#include <iostream>
#include <thread>
#include <mutex>
#include <deque>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/un.h>
#include <netinet/in.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <arpa/inet.h>
using namespace std;

typedef unsigned char BYTE

// Thread safe queuing
mutex queMu;
deque<BYTE*> callQueue;

void add(BYTE* bytes) {
    lock_guard<mutex> locker(queMu);
    callQueue.push_back(bytes);
}
deque<BYTE*> getAll() {
    lock_guard<mutex> locker(queMu);
    deque<BYTE*> ret = callQueue; // deep copy
    callQueue.clear();
    return ret;
}

// listener listens for messages from the main server
void listener(int sockfd) {
    int size = 0;
    const int buffSize = 1400;
    BYTE recvBuff[buffSize];
    memset(recvBuff, '0' ,sizeof(recvBuff));

    while ((n = read(sockfd, recvBuff, buffSize-1) > 0) {
        BYTE* data = new BYTE[n];
        memcpy(data, recvBuff, n);
        add(data);
    }
    exit(1);
}

int connect(string USDAddr) {
    int sockfd = 0;
    struct sockaddr_un serv_addr;

    if ((sockfd = socket(AF_UNIX, SOCK_STREAM, 0))< 0) {
        cout << "Could not create socket\n";
        exit(1);
    }

    memcpy(serv_addr.sun_path, UDSaddr.c_str(), UDSaddr.length());
    serv_addr.sun_family = AF_UNIX;
    serv_addr.sun_path[UDSaddr.length()] = '\0';

    if (connect(sockfd, (struct sockaddr *)&serv_addr, sizeof(serv_addr))<0) {
        cout << "Connection Failed\n";
        exit(1);
    }

    return sockfd;
}

int main(int argc, char const *argv[]) {
    string USDAddr;
    int TickRate;
    if (argc != 3) {
        cout << "Usage: " << argv[0] << " <UDS Address> <Tick Rate>\n";
        return 1;
    }
    USDAddr = argv[1];
    TickRate = argv[2];
    if (TickRate < 0) {
        cout << "Tick rate must be above 0.\n";
        return 1;
    }

    int sockfd = connect(USDAddr);

    thread lis(listener, sockfd);
    lis.detach();

    return 0;
}
