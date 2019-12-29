#include "n2n/n2n.h"

// See https://github.com/pojntfx/n2n/blob/dev/example_sn_embed.c
int supernode_start(int lport, int mgmt_port)
{
    int keep_running;
    n2n_sn_t sss_node;
    int rc;

    sn_init(&sss_node);
    sss_node.daemon = 0;
    sss_node.lport = lport;

    sss_node.sock = open_socket(sss_node.lport, 1);
    if (-1 == sss_node.sock)
    {
        exit(-2);
    }

    sss_node.mgmt_sock = open_socket(mgmt_port, 0);
    if (-1 == sss_node.mgmt_sock)
    {
        exit(-2);
    }

    keep_running = 1;
    rc = run_sn_loop(&sss_node, &keep_running);

    sn_term(&sss_node);

    return rc;
}