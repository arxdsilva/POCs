# Bambu Lab X1C Printing Guide: Microgreens Trays

## Quick Summary

- **Material**: PETG
- **Support**: Tree supports (recommended) or linear supports
- **Layer Height**: 0.2 mm (standard) or 0.16 mm (higher detail)
- **Wall Thickness**: 3–4 walls minimum
- **Brim**: Enabled
- **Bed Temperature**: 70–80°C
- **Nozzle Temperature**: 230–240°C
- **Print Speed**: Default or reduced (consider ~95–100% for better quality)

---

## Detailed Configuration

### Material & Temperature

**PETG Settings** (most reliable for functional trays):
- Nozzle: 235–240°C
- Bed: 75°C
- Chamber (if available): 35°C helps with warp resistance
- Cooling: Default (fan ON after first few layers)

Alternative: **ASA** if you want outdoor UV resistance for the water tray
- Nozzle: 240–250°C
- Bed: 80–100°C
- No cooling (turn off fans)

### Print Quality & Strength

| Setting | Value | Notes |
|---------|-------|-------|
| **Layer Height** | 0.2 mm | Standard; provides good strength-to-time balance. Use 0.16 mm for finer details in drainage holes. |
| **Wall Count** | 4 | Use 4 for maximum durability. Minimum 3 for cost reduction. |
| **Top/Bottom Layers** | 4–5 | Important for watertightness; at least 4 layers. |
| **Infill** | 15–20% | Honeycomb or gyroid pattern. More than 20% adds weight without much benefit. |
| **Line Width** | Default (0.4 mm) | No adjustment needed. |

### Support Settings

#### **RECOMMENDED: No Supports (Drain Holes Facing DOWN)**
- **Orientation**: Place grow tray with drainage holes facing DOWN, flat on build plate
- **Support**: **Disabled**
- **Reason**: Avoids the complexity/memory issues with supporting dense drainage hole patterns
  - Linear supports cause G-code path overflow errors
  - Tree supports exhaust memory on dense hole grids
- **Result**: Fastest, most reliable print with no support removal

#### **Alternative: Tree Supports (if drainage holes face UP)**
- If you must have holes facing up:
  - **Support Type**: Tree (not linear)
  - **Support Angle**: 50–60° (steeper angle = fewer supports)
  - **Z Gap**: 0.3 mm (allows easier removal; single-layer supports often fail)
  - **XY Gap**: 0.3 mm
  - **Support on Build Plate Only**: **Enabled** — prevents support inside hollow areas
  - **Density**: **50–70%** (lower density reduces memory; still provides adequate strength)
  - **Pattern**: Rectilinear
- **Before slicing**: Reduce drainage hole density in script parameters:
  - Increase `DRAIN_HOLE_SPACING_MM` from 15 to 20–25 mm
  - Increase `DRAIN_HOLE_EDGE_MARGIN_MM` from 14 to 20–25 mm
  - This reduces hole count from ~16 to ~9, dramatically lowering support complexity

#### **NOT RECOMMENDED: Linear Supports**
- Linear supports fail on dense drainage patterns (G-code exceeds plate boundaries)
- Use tree or no supports instead

### Bed Adhesion

| Setting | Value | Reason |
|---------|-------|--------|
| **Brim** | Enabled (8–10 mm width) | Microgreens trays are fairly flat and lightweight; brim prevents edge lift. |
| **Raft** | Disabled | Not needed with brim; adds waste material. |
| **First Layer Speed** | 80–100% | Standard Bambu defaults are fine. |
| **Build Plate Type** | Textured (default) | Works well for PETG. Clean plate thoroughly before print. |

---

## Orientation on Print Bed

### Grow Tray (with drainage holes)
**Option A — Drainage holes facing DOWN** (RECOMMENDED — no supports needed)
- Tray sits flat on bed
- No overhangs; no supports required
- Fastest print, zero support waste
- Drainage holes have minor texture from bed contact, but fully functional
- Use this to avoid linear/tree support issues

**Option B — Drainage holes facing UP** (requires supports)
- Drainage holes face up; cleaner visual appearance
- Requires tree supports with 50–70% density (see support section)
- Make sure to reduce drainage hole density in script first to avoid memory/overflow errors
- Longer print time; support removal required

### Water Tray (solid)
- **Orient flat** (solid bottom down or up, either works)
- No overhangs; no supports needed
- Fastest and cleanest print

---

## Step-by-Step Bambu Lab Studio Setup

1. **Import & Orient**
   - Open both STL files in Bambu Lab Studio
   - Orient trays flat on build plate
   - Position grow tray on left half, water tray on right half (to avoid crowding)

2. **Material & Temps**
   - Select filament: PETG
   - Confirm nozzle: 235–240°C
   - Confirm bed: 75°C

3. **Print Profile**
   - Use **Bambu Lab PETG - Standard** profile as base
   - Set **Layer Height**: 0.2 mm
   - Set **Walls**: 4
   - Set **Infill**: 15–20% (honeycomb)

4. **Support**
   - If drainage holes face up: **Enable tree supports** with "support on build plate only" checked
   - If drainage holes face down: **Disable supports**

5. **Adhesion**
   - Enable **Brim** (8 mm)
   - Disable **Raft**

6. **Preview & Slice**
   - Review supports in preview (should be minimal, only on external areas)
   - Check estimated time and weight
   - Slice and send to printer

---

## Estimated Print Times & Material

**Per Tray** (0.2 mm layers, 4 walls, 15% infill):
| Tray | Time | Weight | Notes |
|------|------|--------|-------|
| Grow (with supports) | ~3–4 hours | 45–55 g | Includes support material |
| Grow (no supports) | ~2–3 hours | 40–45 g | If printed holes-down |
| Water (no supports) | ~2–3 hours | 50–60 g | Solid bottom adds weight |

**Both together**: ~5–7 hours on X1C (depending on spacing and support density)

---

## Post-Print Cleanup

1. **Remove from Bed**
   - Let cool to room temperature (~30 min)
   - Flex build plate gently to pop off prints

2. **Clean Supports**
   - Use support-removal tool (included with X1C) or small pliers
   - For tree supports: snap from base; usually clean separation
   - For internal supports in grow tray: carefully extract with tweezers

3. **Sand (Optional)**
   - Light sand (220–400 grit) on top surfaces for a polished finish
   - Focus on any rough edges or support scars

4. **Inspect**
   - Check all drainage holes are clear (3.5 mm diameter)
   - Verify no layer shifts or defects

---

## Troubleshooting

| Issue | Solution |
|-------|----------|
| **Linear supports → "G-code path beyond plate boundaries" error** | Use tree supports with 50–70% density instead. Or rotate tray so drainage holes face DOWN (no supports needed). Reduce drainage hole count by increasing `DRAIN_HOLE_SPACING_MM` to 20–25 mm in script. |
| **Tree supports → Program closes/out of memory** | Reduce drainage hole density in script BEFORE exporting STL: increase `DRAIN_HOLE_SPACING_MM` to 20–25 mm and `DRAIN_HOLE_EDGE_MARGIN_MM` to 20–25 mm. This cuts hole count in half. Re-export STL and retry slicing. |
| Trays warp at edges | Increase bed temp to 80°C, enable brim, reduce nozzle temp by 5°C |
| Drainage holes clogged | Use 0.16 mm layer height or increase hole diameter in script to 4.0 mm |
| Supports hard to remove | Check "support on build plate only" is enabled; use higher Z gap (0.3 mm) |
| First layer adhesion poor | Clean build plate with IPA, ensure bed is level, increase first-layer speed to 110% |
| Watertightness issues | Add 1–2 extra top/bottom layers; ensure water tray bottom is ≥3.0 mm thick |

---

## Final Notes

- **Print both trays in a single run** for convenience (they fit on X1C bed)
- **PETG is hydrophobic** — water will bead on surfaces, but seams may still weep slightly. If perfect watertightness is critical, consider **ASA** or apply food-safe waterproof sealant to bottom seams
- **Test fit** after printing: stack grow tray inside water tray to verify clearances before use
- Keep **support density at 100%** to ensure clean, strong prints

Enjoy your microgreens setup!
