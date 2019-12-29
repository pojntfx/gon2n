#include "n2n/n2n.h"

int supernode_start(int listenPort, int managementPort)
{
    static n2n_sn_t sss_node;
    static int keep_running;

    sn_init(&sss_node);

    sss_node.lport = listenPort;
    sss_node.daemon = 0;

    sss_node.sock = open_socket(sss_node.lport, 1);
    if (-1 == sss_node.sock)
    {
        exit(-2);
    }

    sss_node.mgmt_sock = open_socket(managementPort, 0);
    if (-1 == sss_node.mgmt_sock)
    {
        exit(-2);
    }

    keep_running = 1;
    return run_sn_loop(&sss_node, &keep_running);
}