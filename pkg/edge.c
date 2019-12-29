#include "n2n/n2n.h"

// See https://github.com/pojntfx/n2n/blob/dev/example_edge_embed.c
int edge_start(
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
    int transop_id,
    char *dev,
    char *address_mode,
    char *device_ip,
    char *device_mask,
    char *device_mac,
    int mtu)
{
    int keep_running;
    n2n_edge_conf_t conf;
    tuntap_dev tuntap;
    n2n_edge_t *eee;
    int rc;

    edge_init_conf_defaults(&conf);
    conf.allow_p2p = allow_p2p;
    conf.allow_routing = allow_routing;
    snprintf((char *)conf.community_name, sizeof(conf.community_name), "%s", community_name);
    conf.disable_pmtu_discovery = disable_pmtu_discovery;
    conf.drop_multicast = drop_multicast;
    conf.dyn_ip_mode = dyn_ip_mode;
    conf.encrypt_key = encrypt_key;
    conf.local_port = local_port;
    conf.mgmt_port = mgmt_port;
    conf.register_interval = register_interval;
    conf.register_ttl = register_ttl;
    edge_conf_add_supernode(&conf, supernode);
    conf.tos = tos;
    conf.transop_id = transop_id;

    if (edge_verify_conf(&conf) != 0)
    {
        return -1;
    }

    if (tuntap_open(&tuntap,
                    dev,
                    address_mode,
                    device_ip,
                    device_mask,
                    device_mac,
                    mtu) < 0)
    {
        return -1;
    }

    eee = edge_init(&tuntap, &conf, &rc);
    if (eee == NULL)
    {
        exit(1);
    }

    keep_running = 1;
    rc = run_edge_loop(eee, &keep_running);

    edge_term(eee);
    tuntap_close(&tuntap);

    return rc;
}
