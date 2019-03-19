"""
IP Address utility to get bits, subnets, netmasks, IP addresses, etc.
pylint: 9.91/10
radon: 51.92/100
"""
import os, re, ipaddress, multiprocessing

def is_valid(ip_netmask):
    """
    Check if the string or number is a valid IP address or netmask.
    @param ip_netmask string or integer, e.g. 172.31.146.28, 255.255.255.0, 24.
    @return True | False, whether input is valid
    """
    if type(ip_netmask) is str and "." in ip_netmask:
        octets = re.split(r'\.', ip_netmask)
        if len(octets) is not 4:
            print "Incorrect netmask", ip_netmask
        elif False in [octet.isdigit() for octet in octets]:
            print "Invalid octets", ip_netmask
        elif False in [0 <= int(octet) <= 255 for octet in octets]:
            print "%s is invalid. Should be between 0 and 255." % ip_netmask
        else:
            return True
    elif type(ip_netmask) is str and ip_netmask.isdigit()\
            and 0 <= int(ip_netmask) <= 32:
        return True
    elif type(ip_netmask) is int and 0 <= ip_netmask <= 32:
        return True
    else:
        if type(ip_netmask) is str:
            print ip_netmask + " is invalid."
        elif type(ip_netmask) is int:
            print "%s is invalid." % ip_netmask

    return False

def convert_cidr_to_netmask(cidr):
    """
    Converts a classless inter-domain routing number to netmask.
    @param cidr [string|int] 0-32
    return string or None if invalid
    """
    if type(cidr) is str and cidr.isdigit():
        return convert_cidr_to_netmask(int(cidr))
    elif type(cidr) is int:
        if 0 <= int(cidr) <= 32:
            # Just trying to improve my pylint score ;)
            netmask = [octet * 0 for octet in range(4)]
            remaining_bits = cidr
            index = 0
            while remaining_bits > 0:
                if remaining_bits >= 8:
                    cidr_bits_used = 8
                    netmask[index] = 255
                else:
                    cidr_bits_used = remaining_bits
                    octet = int('1' * cidr_bits_used, 2) << 8 - cidr_bits_used
                    netmask[index] = octet

                remaining_bits -= cidr_bits_used
                index += 1

            return '.'.join([str(octet) for octet in netmask])

        else:
            print "Invalid CIDR", cidr
            return None
    else:
        print "Invalid CIDR", cidr
        return None

def get_binary_bits(cidr_ip_netmask):
    """
    @param cidr_ip_netmask [string|int] classless inter-domain routing,
           e.g. 24; ipv4 address; netmask
    @return list of binary bit
    """
    if is_valid(cidr_ip_netmask):
        if type(cidr_ip_netmask) is str and '.' in cidr_ip_netmask:
            return [format(int(octet), '08b')\
                    for octet in re.split(r'\.', cidr_ip_netmask)]
        else:
            ip_netmask = convert_cidr_to_netmask(cidr_ip_netmask)
            return get_binary_bits(ip_netmask)
    else:
        print "Invalid input", cidr_ip_netmask

def get_network_address(ip_address, netmask):
    """
    Get the network address.
    @param ip_address [string]
    @param netmask [string]
    """
    if is_valid(ip_address) and is_valid(netmask):
        if type(netmask) is str and '.' in netmask:
            return  '.'.join([str(int(re.split(r'\.', ip_address)[i])\
                    & int(re.split(r'\.', netmask)[i])) for i in range(4)])
        else:
            netmask_str = convert_cidr_to_netmask(netmask)
            return  '.'.join([str(int(re.split(r'\.', ip_address)[i])\
                    & int(re.split(r'\.', netmask_str)[i])) for i in range(4)])

def get_host_bits(netmask):
    """
    Get the host bits of a subnet.
    @param netmask, e.g. 255.255.255.0 or CIDR e.g. 24.
    @return Integer
    """
    if is_valid(netmask) and type(netmask):
        # Find index for end of subnet bits
        binary_netmask = get_binary_bits(netmask)
        list.reverse(binary_netmask)
        reversed_binary_netmask = [octet[::-1] for octet in binary_netmask]
        index = -1
        subnet_octet = 4
        for octet in reversed_binary_netmask:
            try:
                index = octet.index('1')
                break
            except ValueError:
                subnet_octet -= 1

        # Perform calculation
        bits_in_octet = 8
        total_bits = 32
        return total_bits - bits_in_octet * subnet_octet + index
#        if __debug__:
#            print "host_bits(%s) = total_bits(%s) - bits_in_octet(%s)\
#                  * subnet_octet(%s) + index(%s)" % (host_bits, total_bits,\
#                  bits_in_octet, subnet_octet, index)
    elif is_valid(netmask) == 2:
        if type(netmask) is int:
            return 32 - netmask
        else:
            return 32 - int(netmask)

def get_number_of_hosts(netmask):
    """
    Calculate the number of hosts from a netmask.
    @param netmask - 4 octet string or number of bits,
           e.g. 255.255.255.0 or 24.
    @return 0 =< n | -1 for error
    """
    if is_valid(netmask):
        host_bits = get_host_bits(netmask)
        if host_bits:
            return 2**host_bits - 2
    elif is_valid(netmask) == 2:
        # This is ok for <= 32 bits
        return 2**get_host_bits(netmask) - 2

    return -1

def get_broadcast_address(ip_address, netmask):
    """
    Get the broadcast IP address of a subnet.
    @param ip_address Any IP address on the subnet.
    @param netmask
    @return String or None if invalid IP address
    """
    if is_valid(ip_address):
        number_of_host_bits = get_host_bits(netmask)
        host_bits = [0 for i in range(4)]
        remaining_host_bits = number_of_host_bits
        index = 3
        while remaining_host_bits > 0:
            if remaining_host_bits >= 8:
                host_bits_used = 8
            else:
                host_bits_used = remaining_host_bits

            remaining_host_bits -= host_bits_used

            host_octet_binary = format(int('1' * host_bits_used, 2), '08b')
            host_bits[index] |= int(host_octet_binary, 2)
            index -= 1

        return '.'.join([str(int(re.split(r'\.', ip_address)[i])\
                | host_bits[i]) for i in range(4)])
    return None

def get_ip_list(ip_address, netmask):
    """
    Get the list of host IP addresses in a subnet
    excluding broadcast and network addresses.
    @param ip_address [string]
    @param netmask [string] E.g. 255.255.255.0 or CIDR
    """
    ip_list = list()
    if is_valid(ip_address) and is_valid(netmask):
        network_address = ipaddress.IPv4Address\
                (get_network_address(ip_address, netmask))
        number_of_hosts = get_number_of_hosts(netmask)
        for i in range(1, number_of_hosts + 1):
            ip_list.append(str(network_address + i))
    else:
        print "Please check IP address or netmask."
    return ip_list


def ping_sweep(ip_list, is_alive=True, detail=False):
    """
    Pings a list of IP addresses.
    @param ip_list [string] list of IP addresses.
    @param is_alive [boolean] True for responding, False for available.
    @param detail [boolean] Output process details.
    @return list of IP addresses
    """

    def ping(ip_q, results_q):
        """
        Ping a subnet.
        @param ip_q IP addresses queue
        @param results_q Results queue
        """

        if detail:
            ping_cmd = "ping -c 1 %s"
        else:
            ping_cmd = "ping -c 1 %s >> /dev/null"

        while True:
            ip_address = ip_q.get()
            if ip_address is None:
                break

            try:
                results_q.put\
                    ({ip_address : os.system(ping_cmd % ip_address)})
            except:
                pass
    # end ping

    return_ip_list = []
    if isinstance(ip_list, list):

        ip_queue = multiprocessing.Queue()
        for ip_addr in ip_list:
            ip_queue.put(ip_addr)

        results_queue = multiprocessing.Queue()
        ip_pool = [multiprocessing.Process(target=ping,\
                args=(ip_queue, results_queue)) for i in range(len(ip_list))]

        for ip_addr in ip_pool:
            ip_addr.start()

        for ip_addr in ip_pool:
            ip_queue.put(None)

        for ip_addr in ip_pool:
            ip_addr.join()

        while not results_queue.empty():
            result = results_queue.get()
            for ip_addr, status in result.items():
                if is_alive:
                    if status == 0:
                        return_ip_list.append(ip_addr)
                else:
                    if status != 0:
                        return_ip_list.append(ip_addr)
    else:
        print "Input argument must be a list of IP addresses."
    return return_ip_list
# end ping_sweep
