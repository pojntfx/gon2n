#include "n2n/n2n.h"

#define N2N_SN_LPORT_DEFAULT 7654
#define N2N_SN_MGMT_PORT 5645

typedef struct sn_stats
{
    size_t errors;         /* Number of errors encountered. */
    size_t reg_super;      /* Number of REGISTER_SUPER requests received. */
    size_t reg_super_nak;  /* Number of REGISTER_SUPER requests declined. */
    size_t fwd;            /* Number of messages forwarded. */
    size_t broadcast;      /* Number of messages broadcast to a community. */
    time_t last_fwd;       /* Time when last message was forwarded. */
    time_t last_reg_super; /* Time when last REGISTER_SUPER was received. */
} sn_stats_t;

typedef struct n2n_sn
{
    time_t start_time; /* Used to measure uptime. */
    sn_stats_t stats;
    int daemon;           /* If non-zero then daemonise. */
    uint16_t lport;       /* Local UDP port to bind to. */
    int sock;             /* Main socket for UDP traffic with edges. */
    int mgmt_sock;        /* management socket. */
    int lock_communities; /* If true, only loaded communities can be used. */
    struct sn_community *communities;
} n2n_sn_t;

int supernode_start()
{
    static n2n_sn_t sss_node;

    memset(&sss_node, 0, sizeof(n2n_sn_t));

    // TODO: Set these struct properties
    sss_node.daemon = 0;                   // Don't daemonise
    sss_node.lock_communities = 0;         // Allow all, not just the loaded communities
    sss_node.lport = N2N_SN_LPORT_DEFAULT; // Lsten for UDP traffic on port 1234

    // Listen on UDP (main)
    sss_node.sock = open_socket(sss_node.lport, 1 /*bind ANY*/);

    // Listen on UDP (management)
    sss_node.mgmt_sock = open_socket(N2N_SN_MGMT_PORT, 0 /* bind LOOPBACK */);

    // TODO: Handle signals

    return 1;
}