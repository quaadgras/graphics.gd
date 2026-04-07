#include <gd.h>

typedef Object Node_t;

struct {
    Object (*new)();
    void (*set_name)(Node_t node, StringName name);
    void (*queue_free)(Node_t node);
    void (*free)(Node_t node);
} Node;
