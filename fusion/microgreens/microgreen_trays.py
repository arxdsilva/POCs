# Fusion 360 Parametric Microgreens Tray Generator
# Creates two printable trays for a Bambu Lab X1C:
# 1) Grow tray with drainage holes
# 2) Water tray with solid bottom
#
# How to use:
# 1. Open Fusion 360
# 2. Go to Utilities > Scripts and Add-Ins
# 3. Create a new Python script
# 4. Replace its contents with this file
# 5. Run the script
# 6. Edit values in the PARAMETERS section if needed

import adsk.core
import adsk.fusion
import adsk.cam
import traceback
import math


# =========================
# PARAMETERS - EDIT THESE
# =========================

# Bambu Lab X1C usable safe footprint is slightly below 256 x 256 mm.
# Keep water tray near max print area; grow tray is auto-sized to nest inside.
WATER_TRAY_OUTER_WIDTH_MM = 248.0
WATER_TRAY_OUTER_DEPTH_MM = 248.0

# X1C bed constraints.
# Some X1C setups have a bottom-left exclusion rectangle that the toolhead cannot reach.
# These values apply a conservative centered-placement limit to avoid that region.
X1C_BED_WIDTH_MM = 256.0
X1C_BED_DEPTH_MM = 256.0
X1C_EDGE_CLEARANCE_MM = 2.0
X1C_BOTTOM_LEFT_EXCLUSION_WIDTH_MM = 18.0
X1C_BOTTOM_LEFT_EXCLUSION_DEPTH_MM = 18.0

# Grow tray
GROW_TRAY_HEIGHT_MM = 40.0
GROW_TRAY_WALL_THICKNESS_MM = 3.0
GROW_TRAY_BOTTOM_THICKNESS_MM = 2.4

# Water tray
WATER_TRAY_HEIGHT_MM = 30.0
WATER_TRAY_WALL_THICKNESS_MM = 3.0
WATER_TRAY_BOTTOM_THICKNESS_MM = 3.0

# Nesting clearance between grow tray outer walls and water tray inner walls.
# This is per-side clearance. Example: 1.2 mm means ~2.4 mm total width gap.
STACK_SIDE_CLEARANCE_MM = 1.2

# Corner radius
OUTER_CORNER_RADIUS_MM = 10.0
INNER_CORNER_RADIUS_MM = 7.0

# Drainage holes for grow tray
DRAIN_HOLE_DIAMETER_MM = 3.5
DRAIN_HOLE_SPACING_MM = 15.0
DRAIN_HOLE_EDGE_MARGIN_MM = 14.0

# Feet / raised ribs under grow tray
ADD_RAISED_SUPPORT_RIBS = True
RIB_WIDTH_MM = 5.0
RIB_HEIGHT_MM = 3.0
RIB_COUNT_EACH_DIRECTION = 3

# Tray spacing in Fusion workspace
COMPONENT_SPACING_MM = 285.0

# Optional labels embossed into the side
ADD_LABELS = True
TEXT_HEIGHT_MM = 8.0
TEXT_DEPTH_MM = 0.6


def get_grow_outer_width_mm():
    return WATER_TRAY_OUTER_WIDTH_MM - 2.0 * WATER_TRAY_WALL_THICKNESS_MM - 2.0 * STACK_SIDE_CLEARANCE_MM


def get_grow_outer_depth_mm():
    return WATER_TRAY_OUTER_DEPTH_MM - 2.0 * WATER_TRAY_WALL_THICKNESS_MM - 2.0 * STACK_SIDE_CLEARANCE_MM


def fit_to_x1c_reachable_area_mm(requested_width_mm, requested_depth_mm):
    """Fits a rectangle to X1C reachable area with an optional bottom-left exclusion zone.

    The exclusion is a single corner rectangle. A part is printable if it can be shifted
    so that either its left edge is to the right of the exclusion or its front edge is
    above the exclusion.

    Returns: (fit_width_mm, fit_depth_mm, was_adjusted)
    """
    bed_w = max(1.0, X1C_BED_WIDTH_MM - 2.0 * X1C_EDGE_CLEARANCE_MM)
    bed_d = max(1.0, X1C_BED_DEPTH_MM - 2.0 * X1C_EDGE_CLEARANCE_MM)

    fit_w = min(requested_width_mm, bed_w)
    fit_d = min(requested_depth_mm, bed_d)
    adjusted = (fit_w != requested_width_mm) or (fit_d != requested_depth_mm)

    excl_w = max(0.0, X1C_BOTTOM_LEFT_EXCLUSION_WIDTH_MM)
    excl_d = max(0.0, X1C_BOTTOM_LEFT_EXCLUSION_DEPTH_MM)

    # No corner exclusion configured.
    if excl_w <= 0.0 or excl_d <= 0.0:
        return fit_w, fit_d, adjusted

    # Feasible placement exists if we can shift in X or in Y enough to clear the corner.
    can_shift_x = (bed_w - fit_w) >= excl_w
    can_shift_y = (bed_d - fit_d) >= excl_d
    if can_shift_x or can_shift_y:
        return fit_w, fit_d, adjusted

    # Not placeable as-is: shrink the axis that needs the least reduction.
    target_w = max(1.0, bed_w - excl_w)
    target_d = max(1.0, bed_d - excl_d)
    reduce_w = max(0.0, fit_w - target_w)
    reduce_d = max(0.0, fit_d - target_d)

    if reduce_w <= reduce_d:
        fit_w = target_w
    else:
        fit_d = target_d

    return fit_w, fit_d, True


# =========================
# HELPERS
# =========================

def mm(value):
    return value / 10.0  # Fusion API internal unit is cm


def create_centered_rounded_rect(sketch, width_mm, depth_mm, radius_mm, center_x_mm=0.0):
    """Creates a centered rounded rectangle profile using lines and arcs."""
    lines = sketch.sketchCurves.sketchLines
    arcs = sketch.sketchCurves.sketchArcs

    w = mm(width_mm)
    d = mm(depth_mm)
    r = mm(radius_mm)

    cx = mm(center_x_mm)
    x = w / 2.0
    y = d / 2.0

    # Points clockwise, with tangent line endpoints
    p1 = adsk.core.Point3D.create(cx - x + r, -y, 0)
    p2 = adsk.core.Point3D.create(cx + x - r, -y, 0)
    p3 = adsk.core.Point3D.create(cx + x, -y + r, 0)
    p4 = adsk.core.Point3D.create(cx + x, y - r, 0)
    p5 = adsk.core.Point3D.create(cx + x - r, y, 0)
    p6 = adsk.core.Point3D.create(cx - x + r, y, 0)
    p7 = adsk.core.Point3D.create(cx - x, y - r, 0)
    p8 = adsk.core.Point3D.create(cx - x, -y + r, 0)

    lines.addByTwoPoints(p1, p2)
    arcs.addByCenterStartSweep(
        adsk.core.Point3D.create(cx + x - r, -y + r, 0),
        p2,
        math.radians(90)
    )
    lines.addByTwoPoints(p3, p4)
    arcs.addByCenterStartSweep(
        adsk.core.Point3D.create(cx + x - r, y - r, 0),
        p4,
        math.radians(90)
    )
    lines.addByTwoPoints(p5, p6)
    arcs.addByCenterStartSweep(
        adsk.core.Point3D.create(cx - x + r, y - r, 0),
        p6,
        math.radians(90)
    )
    lines.addByTwoPoints(p7, p8)
    arcs.addByCenterStartSweep(
        adsk.core.Point3D.create(cx - x + r, -y + r, 0),
        p8,
        math.radians(90)
    )


def extrude_profile(root_comp, profile, distance_mm, operation):
    extrudes = root_comp.features.extrudeFeatures
    ext_input = extrudes.createInput(profile, operation)
    ext_input.setDistanceExtent(False, adsk.core.ValueInput.createByReal(mm(distance_mm)))
    return extrudes.add(ext_input)


def create_box_shell(root_comp, name, x_offset_mm, width_mm, depth_mm, height_mm, wall_mm, bottom_mm):
    """Creates a tray shell using outer extrusion and inner cut.

    Returns: (component_used_for_features, sketch_origin_x_mm)
    """
    comp = None
    origin_x_mm = 0.0

    # Assembly docs allow separate components. Part docs do not, so we fall back
    # to building directly in the root component at an x-offset.
    try:
        occs = root_comp.occurrences
        matrix = adsk.core.Matrix3D.create()
        matrix.translation = adsk.core.Vector3D.create(mm(x_offset_mm), 0, 0)
        occ = occs.addNewComponent(matrix)
        comp = occ.component
        comp.name = name
    except RuntimeError:
        comp = root_comp
        origin_x_mm = x_offset_mm

    xy = comp.xYConstructionPlane

    outer_sketch = comp.sketches.add(xy)
    outer_sketch.name = f"{name} outer footprint"
    create_centered_rounded_rect(outer_sketch, width_mm, depth_mm, OUTER_CORNER_RADIUS_MM, origin_x_mm)
    outer_body = extrude_profile(comp, outer_sketch.profiles.item(0), height_mm, adsk.fusion.FeatureOperations.NewBodyFeatureOperation)
    outer_body.bodies.item(0).name = name

    inner_width = width_mm - 2 * wall_mm
    inner_depth = depth_mm - 2 * wall_mm
    inner_sketch = comp.sketches.add(xy)
    inner_sketch.name = f"{name} inner hollow cut"
    create_centered_rounded_rect(inner_sketch, inner_width, inner_depth, INNER_CORNER_RADIUS_MM, origin_x_mm)

    cut_distance = height_mm - bottom_mm
    cut_input = comp.features.extrudeFeatures.createInput(inner_sketch.profiles.item(0), adsk.fusion.FeatureOperations.CutFeatureOperation)
    cut_input.setOneSideExtent(
        adsk.fusion.DistanceExtentDefinition.create(adsk.core.ValueInput.createByReal(mm(cut_distance))),
        adsk.fusion.ExtentDirections.PositiveExtentDirection
    )
    # Start the hollow cut at bottom thickness height
    offset_plane_input = comp.constructionPlanes.createInput()
    offset_plane_input.setByOffset(xy, adsk.core.ValueInput.createByReal(mm(bottom_mm)))
    start_plane = comp.constructionPlanes.add(offset_plane_input)
    cut_input.startExtent = adsk.fusion.FromEntityStartDefinition.create(start_plane, adsk.core.ValueInput.createByReal(0))
    comp.features.extrudeFeatures.add(cut_input)

    return comp, origin_x_mm


def add_drainage_holes(comp, width_mm, depth_mm, bottom_mm, origin_x_mm=0.0):
    """Cuts a grid of circular drainage holes through the grow tray bottom."""
    xy = comp.xYConstructionPlane
    sketch = comp.sketches.add(xy)
    sketch.name = "drainage hole pattern"

    circles = sketch.sketchCurves.sketchCircles
    radius = mm(DRAIN_HOLE_DIAMETER_MM / 2.0)

    usable_w = width_mm - 2 * DRAIN_HOLE_EDGE_MARGIN_MM
    usable_d = depth_mm - 2 * DRAIN_HOLE_EDGE_MARGIN_MM
    cols = int(usable_w // DRAIN_HOLE_SPACING_MM) + 1
    rows = int(usable_d // DRAIN_HOLE_SPACING_MM) + 1

    start_x = -((cols - 1) * DRAIN_HOLE_SPACING_MM) / 2.0
    start_y = -((rows - 1) * DRAIN_HOLE_SPACING_MM) / 2.0

    for i in range(cols):
        for j in range(rows):
            x = start_x + i * DRAIN_HOLE_SPACING_MM
            y = start_y + j * DRAIN_HOLE_SPACING_MM
            circles.addByCenterRadius(adsk.core.Point3D.create(mm(origin_x_mm + x), mm(y), 0), radius)

    profiles = adsk.core.ObjectCollection.create()
    for p in sketch.profiles:
        profiles.add(p)

    ext_input = comp.features.extrudeFeatures.createInput(profiles, adsk.fusion.FeatureOperations.CutFeatureOperation)
    ext_input.setDistanceExtent(False, adsk.core.ValueInput.createByReal(mm(bottom_mm + 1.0)))
    comp.features.extrudeFeatures.add(ext_input)

    # If underside ribs exist, continue each hole downward through rib material too.
    if ADD_RAISED_SUPPORT_RIBS and RIB_HEIGHT_MM > 0:
        rib_cut_input = comp.features.extrudeFeatures.createInput(profiles, adsk.fusion.FeatureOperations.CutFeatureOperation)
        rib_cut_input.setDistanceExtent(False, adsk.core.ValueInput.createByReal(mm(-(RIB_HEIGHT_MM + 0.5))))
        comp.features.extrudeFeatures.add(rib_cut_input)


def add_support_ribs(comp, width_mm, depth_mm, origin_x_mm=0.0):
    """Adds shallow ribs under grow tray to keep it slightly raised above the water tray."""
    if not ADD_RAISED_SUPPORT_RIBS:
        return

    xy = comp.xYConstructionPlane
    sketch = comp.sketches.add(xy)
    sketch.name = "underside support ribs"

    inner_w = width_mm - 40.0
    inner_d = depth_mm - 40.0
    spacing_x = inner_w / (RIB_COUNT_EACH_DIRECTION + 1)
    spacing_y = inner_d / (RIB_COUNT_EACH_DIRECTION + 1)

    # Horizontal ribs
    for i in range(RIB_COUNT_EACH_DIRECTION):
        y = -inner_d / 2 + spacing_y * (i + 1)
        sketch.sketchCurves.sketchLines.addCenterPointRectangle(
            adsk.core.Point3D.create(mm(origin_x_mm), mm(y), 0),
            adsk.core.Point3D.create(mm(origin_x_mm + inner_w / 2), mm(y + RIB_WIDTH_MM / 2), 0)
        )

    # Vertical ribs
    for i in range(RIB_COUNT_EACH_DIRECTION):
        x = -inner_w / 2 + spacing_x * (i + 1)
        sketch.sketchCurves.sketchLines.addCenterPointRectangle(
            adsk.core.Point3D.create(mm(origin_x_mm + x), 0, 0),
            adsk.core.Point3D.create(mm(origin_x_mm + x + RIB_WIDTH_MM / 2), mm(inner_d / 2), 0)
        )

    profiles = adsk.core.ObjectCollection.create()
    for p in sketch.profiles:
        profiles.add(p)

    ext_input = comp.features.extrudeFeatures.createInput(profiles, adsk.fusion.FeatureOperations.JoinFeatureOperation)
    ext_input.setDistanceExtent(False, adsk.core.ValueInput.createByReal(mm(-RIB_HEIGHT_MM)))
    comp.features.extrudeFeatures.add(ext_input)


def add_label(comp, label, width_mm, depth_mm, height_mm, origin_x_mm=0.0):
    if not ADD_LABELS:
        return

    # Put a small debossed-looking label on the front wall by extruding text slightly outward.
    yz_plane = comp.yZConstructionPlane
    offset_plane_input = comp.constructionPlanes.createInput()
    offset_plane_input.setByOffset(yz_plane, adsk.core.ValueInput.createByReal(mm(origin_x_mm + width_mm / 2.0 + 0.01)))
    plane = comp.constructionPlanes.add(offset_plane_input)

    sketch = comp.sketches.add(plane)
    sketch.name = f"{label} side label"

    # Some Fusion builds expect a raw double for createInput2 height.
    # If text creation fails, we skip labels rather than failing tray generation.
    try:
        text_input = sketch.sketchTexts.createInput2(label, mm(TEXT_HEIGHT_MM))
        text_input.setAsMultiLine(
            adsk.core.Point3D.create(mm(-depth_mm / 2 + 18), mm(height_mm / 2 - 4), 0),
            adsk.core.Point3D.create(mm(depth_mm / 2 - 18), mm(height_mm / 2 + 8), 0),
            adsk.core.HorizontalAlignments.CenterHorizontalAlignment,
            adsk.core.VerticalAlignments.MiddleVerticalAlignment,
            0
        )
        sketch.sketchTexts.add(text_input)

        # Text profiles may not always be available immediately depending on Fusion version.
        # If unavailable, label creation silently skips extrusion.
        profiles = adsk.core.ObjectCollection.create()
        for p in sketch.profiles:
            profiles.add(p)
        if profiles.count > 0:
            ext_input = comp.features.extrudeFeatures.createInput(profiles, adsk.fusion.FeatureOperations.JoinFeatureOperation)
            ext_input.setDistanceExtent(False, adsk.core.ValueInput.createByReal(mm(TEXT_DEPTH_MM)))
            comp.features.extrudeFeatures.add(ext_input)
    except:
        pass


def run(context):
    ui = None
    try:
        app = adsk.core.Application.get()
        ui = app.userInterface
        design = app.activeProduct

        if not isinstance(design, adsk.fusion.Design):
            ui.messageBox("Please open a Fusion 360 design before running this script.")
            return

        root = design.rootComponent

        water_width_mm, water_depth_mm, was_adjusted = fit_to_x1c_reachable_area_mm(
            WATER_TRAY_OUTER_WIDTH_MM,
            WATER_TRAY_OUTER_DEPTH_MM,
        )

        grow_outer_width_mm = water_width_mm - 2.0 * WATER_TRAY_WALL_THICKNESS_MM - 2.0 * STACK_SIDE_CLEARANCE_MM
        grow_outer_depth_mm = water_depth_mm - 2.0 * WATER_TRAY_WALL_THICKNESS_MM - 2.0 * STACK_SIDE_CLEARANCE_MM

        if grow_outer_width_mm <= 2 * GROW_TRAY_WALL_THICKNESS_MM + 10 or grow_outer_depth_mm <= 2 * GROW_TRAY_WALL_THICKNESS_MM + 10:
            ui.messageBox(
                "Invalid sizing: grow tray is too small after clearances.\n"
                "Reduce STACK_SIDE_CLEARANCE_MM or water tray wall thickness."
            )
            return

        grow, grow_origin_x = create_box_shell(
            root,
            "Microgreens Grow Tray - Drainage",
            -COMPONENT_SPACING_MM / 2,
            grow_outer_width_mm,
            grow_outer_depth_mm,
            GROW_TRAY_HEIGHT_MM,
            GROW_TRAY_WALL_THICKNESS_MM,
            GROW_TRAY_BOTTOM_THICKNESS_MM,
        )
        add_support_ribs(grow, grow_outer_width_mm, grow_outer_depth_mm, grow_origin_x)
        add_drainage_holes(grow, grow_outer_width_mm, grow_outer_depth_mm, GROW_TRAY_BOTTOM_THICKNESS_MM, grow_origin_x)
        add_label(grow, "GROW", grow_outer_width_mm, grow_outer_depth_mm, GROW_TRAY_HEIGHT_MM, grow_origin_x)

        water, water_origin_x = create_box_shell(
            root,
            "Microgreens Water Tray - Solid",
            COMPONENT_SPACING_MM / 2,
            water_width_mm,
            water_depth_mm,
            WATER_TRAY_HEIGHT_MM,
            WATER_TRAY_WALL_THICKNESS_MM,
            WATER_TRAY_BOTTOM_THICKNESS_MM,
        )
        add_label(water, "WATER", water_width_mm, water_depth_mm, WATER_TRAY_HEIGHT_MM, water_origin_x)

        if was_adjusted:
            ui.messageBox(
                "Microgreens tray system created with X1C reachability adjustment.\n\n"
                f"Requested water tray: {WATER_TRAY_OUTER_WIDTH_MM:.1f} x {WATER_TRAY_OUTER_DEPTH_MM:.1f} mm\n"
                f"Generated water tray: {water_width_mm:.1f} x {water_depth_mm:.1f} mm\n\n"
                "If this is too small, reduce X1C_BOTTOM_LEFT_EXCLUSION_* to your measured values."
            )
            return

        ui.messageBox(
            "Microgreens tray system created.\n\n"
            "Export each component as STL/3MF for printing.\n"
            "Recommended: PETG, 0.2 mm layer height, 3-4 walls, brim enabled."
        )

    except Exception:
        if ui:
            ui.messageBox("Failed:\n{}".format(traceback.format_exc()))
