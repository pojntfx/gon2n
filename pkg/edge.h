#include "n2n/n2n.h"

int edge_configure(n2n_edge_conf_t *conf,
                   int allow_p2p,
                   int allow_routing,
                   char *community_name,
                   int disable_pmtu_discovery,
                   int drop_multicast,
                   int dyn_ip_mode,
                   char *encrypt_key,
                   int local_port,
                   int mgmt_port,
                   int register_interval,
                   int register_ttl,
                   char *supernode,
                   int tos,
                   int transop_id);

int edge_start(tuntap_dev *tuntap, n2n_edge_conf_t *conf, int *keep_running);