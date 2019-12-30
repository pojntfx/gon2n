#include "n2n/n2n.h"

// See https://github.com/pojntfx/n2n/blob/dev/example_sn_embed.c

int supernode_configure(n2n_sn_t *sss_node, int lport)
{
    const int rc = sn_init(sss_node);
    sss_node->daemon = 0;
    sss_node->lport = lport;

    return rc;
}

int supernode_open_lport_socket(n2n_sn_t *sss_node)
{
    sss_node->sock = open_socket(sss_node->lport, 1);
    if (-1 == sss_node->sock)
    {
        exit(-2);
    }

    return 0;
}

int supernode_open_mgmt_socket(n2n_sn_t *sss_node, int mgmt_port)
{
    sss_node->mgmt_sock = open_socket(mgmt_port, 0);
    if (-1 == sss_node->mgmt_sock)
    {
        exit(-2);
    }

    return 0;
}

int supernode_start(n2n_sn_t *sss_node, int *keep_running)
{
    return run_sn_loop(sss_node, keep_running);
}

int supernode_stop(n2n_sn_t *sss_node)
{
    sn_term(sss_node);

    return 0;
}