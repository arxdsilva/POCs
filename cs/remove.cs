void AVLTree::Delete(int x, TreePointer &p, bool &h) {
    if (p == NULL) {
        count << "Elemento inexistente";
        abort();
    }
    if (x < p->Entry) {
        Delete(x, p->LeftNode, h);
        if (h) { balenceL(p, h); }
        return;
    } 
    if (x > p->Entry) {
        Delete(x, p->RightNode, h);
        if (h) { balanceR(p, h); }
        return;
    }
    TreePointer q;
    q = p;
    if (q->RightNode == NULL) {
        p = q->LeftNode;
        h = true;
        delete q;
        return;
    } 
    if (q->LeftNode == NULL) {
        p = q->RightNode;
        h = true;
    } else {
        DelMin(q, q->RightNode, h);
        if (h) { balanceR(p, h); }
    }
    delete q;
}


void AVLTree::DelMin(TreePointer &q, TreePointer &r, bool &h) {
    if (r->LeftNode != NULL) {
        DelMin(q, r->LeftNode, h);
        if (h) {balanceL(r,h);}
        return;
    }
    q->Entry = r->Entry;
    q->count = r->count;
    q = r;
    r = r->RightNode;
    h = true;
}

