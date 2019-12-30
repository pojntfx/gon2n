#include "n2n/n2n.h"

int supernode_configure(n2n_sn_t *sss_node, int lport);

int supernode_open_lport_socket(n2n_sn_t *sss_node);

int supernode_open_mgmt_socket(n2n_sn_t *sss_node, int mgmt_port);

int supernode_start(n2n_sn_t *sss_node, int *keep_running);

int supernode_stop(n2n_sn_t *sss_node);
