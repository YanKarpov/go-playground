package main

type BaseHandler struct {
    next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
    b.next = handler
    return handler
}

func (b *BaseHandler) UseItem(item string) {
    if b.next != nil {
        b.next.UseItem(item)
    }
}
