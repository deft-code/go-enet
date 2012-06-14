/*
A Go wrapper for the ENet library.

ENetPeer and ENetHost have been wrapped in Go types with methods for the
various ENet functions

ENetEvent has been translated to a Go type.

ENetPacket and ENetAddress have been hidden in the implementation, use []byte
and net.UDPAddr respectively.
*/
package enet
