#include "n2n/n2n.h"

int edge_start(char *device_name, char *community_name,
               char *encrypt_key, char *device_mac,
               char *local_ip_address,
               char *supernode_ip_address_port,
               int *keep_on_running)
{
    return quick_edge_init(device_name, community_name,
                           encrypt_key, device_mac,
                           local_ip_address,
                           supernode_ip_address_port,
                           keep_on_running);
};
