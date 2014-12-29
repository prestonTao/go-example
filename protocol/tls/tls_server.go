/*
	RFC 2818为HTTP Over TLS-网络协议
	TLSv1.2协议
	（注：目前主流的SSL都没有对压缩的支持）

	http://blog.jobbole.com/48369/
	http://www.cnblogs.com/LittleHann/p/3733469.html
	http://blog.csdn.net/sealyao/article/details/5901510
	http://blog.chinaunix.net/uid-20564848-id-74650.html
	http://www.ruanyifeng.com/blog/2014/09/illustration-ssl.html

	http://www.fenesky.com/blog/2014/07/19/how-https-works.html
	PreMaster/Master Secret(Key)计算方式
	http://www.fenesky.com/blog/2014/07/25/how-premaster-secret.html
	Htttps SSL/TLS Session Secret(Key)计算
	http://www.fenesky.com/blog/2014/07/25/how-session-secret.html

	X.509证书解析和使用
	http://blog.csdn.net/sealyao/article/details/5902249
	缓存一致性（Cache Coherency）入门
	http://kb.cnblogs.com/page/504824/


	握手次序使用三个协议：
	1.握手协议，执行客户端和服务器SSL会话的建立过程。
		SSL Handshake Protocol. SSL
	2.更改密码协议，负责协商会话用的密码套接字。
		SSL Change Cipher Spec Protocol .SSL
	3.告警协议，负责在客户端和服务器间传递SSL错误信息。
		SSL Alert Protocol.SSL

	4.应用数据协议


	====================================================================================
	| content type |    major     |   minor   |    length    |  content  |     MAC     |
	====================================================================================
	|   内容类型   |   主要版本   | 次要版本  |  数据包长度  |  记录数据 |     MAC     |
	====================================================================================
	|    1 byte    |   1 byte     |  1 byte   |    2 byte    |  > 0 byte | 0,6 或 20位 |
	====================================================================================

	content type：内容类型(8位)
		1. 握手协议(handshake): 22
		2. 警告协议(alert): 21
		3. 改变密码格式协议(change_cipher_spec): 20
		4. 应用数据协议(application_data): 23

	major: 主要版本(8位)
		使用的SSL主要版本，目前的SSL版本是SSL v3，所以这个字段的值只有3这个值

	minor: 次要版本(8位)
		使用的SSL次要版本。对于SSL v3.1，值为1

	length：数据包长度(16位)
		1) 明文数据包:
		这个字段表示的是明文数据以字节为单位的长度
		2) 压缩数据包
		这个字段表示的是压缩数据以字节为单位的长度
		3) 加密数据包
		这个字段表示的是加密数据以字节为单位的长度

	content: 记录数据
		这个区块封装了上层协议的数据
		  1) 明文数据包:
		  opaque fragment[SSLPlaintext.length];
		  2) 压缩数据包
		  opaque fragment[SSLCompressed.length];
		  3) 加密数据包
		    3.1) 流式(stream)加密: GenericStreamCipher
		      3.1.1) opaque content[SSLCompressed.length];
		      3.1.2) opaque MAC[CipherSpec.hash_size];
		    3.2) 分组(block)加密: GenericBlockCipher
		      3.2.1) opaque content[SSLCompressed.length];
		      3.2.2) opaque MAC[CipherSpec.hash_size];
		      3.2.3) uint8 padding[GenericBlockCipher.padding_length];
		      3.2.4) uint8 padding_length;

	MAC: MAC(0、16、20位)
		MAC进行加密


	## 握手协议

	================================
	| type | length | content |

	1. type: 类型(Type)(1字节)
		该字段指明使用的SSL握手协议报文类型
		    1) hello_request:
		    2) client_hello:
		    3) server_hello:
		    4) certificate:
		    5) server_key_exchange:
		    6) certificate_request:
		    7) server_done:
		    8) certificate_verify:
		    9) client_key_exchange:
		    10) finished:

	2. 长度(Length)(3字节):
		以字节为单位的报文长度。

	3. 内容(Content)(≥1字节):
对应报文类型的的实际内容、参数
    　　1) hello_request: 空
        2) client_hello:
        　　2.1) 版本(ProtocolVersion)
        　　代表客户端可以支持的SSL最高版本号
            　　2.1.1) 主版本: 3
            　　2.1.2) 次版本: 0
        　　2.2) 随机数(Random)
        　　客户端产生的一个用于生成主密钥(master key)的32字节的随机数(主密钥由客户端和服务端的随机数共同生成)
            　　2.2.1) uint32 gmt_unix_time;
            　　2.2.2) opaque random_bytes[28];
        　　4+28=32字节
        　　2.3) 会话ID: opaque SessionID<0..32>;
        　　2.4) 密文族(加密套件):
        　　一个客户端可以支持的密码套件列表。这个列表会根据使用优先顺序排列，每个密码套件都指定了"密钥交换算法(Deffie-Hellman密钥交换算法、基于RSA的密钥交换和另一种实
现在Fortezza chip上的密钥交换)"、"加密算法(DES、RC4、RC2、3DES等)"、"认证算法(MD5或SHA-1)"、"加密方式(流、分组)"
            　　2.4.1) CipherSuite SSL_RSA_WITH_NULL_MD5
            　　2.4.2) CipherSuite SSL_RSA_WITH_NULL_SHA
            　　2.4.3) CipherSuite SSL_RSA_EXPORT_WITH_RC4_40_MD5
            　　2.4.4) CipherSuite SSL_RSA_WITH_RC4_128_MD5
            　　2.4.5) CipherSuite SSL_RSA_WITH_RC4_128_SHA
            　　2.4.6) CipherSuite SSL_RSA_EXPORT_WITH_RC2_CBC_40_MD5
            　　2.4.7) CipherSuite SSL_RSA_WITH_IDEA_CBC_SHA
            　　2.4.8) CipherSuite SSL_RSA_EXPORT_WITH_DES40_CBC_SHA
            　　2.4.9) CipherSuite SSL_RSA_WITH_DES_CBC_SHA
            　　2.4.10) CipherSuite SSL_RSA_WITH_3DES_EDE_CBC_SHA
            　　2.4.11) CipherSuite SSL_DH_DSS_EXPORT_WITH_DES40_CBC_SHA
            　　2.4.12) CipherSuite SSL_DH_DSS_WITH_DES_CBC_SHA
            　　2.4.13) CipherSuite SSL_DH_DSS_WITH_3DES_EDE_CBC_SHA
            　　2.4.14) CipherSuite SSL_DH_RSA_EXPORT_WITH_DES40_CBC_SHA
            　　2.4.15) CipherSuite SSL_DH_RSA_WITH_DES_CBC_SHA
            　　2.4.16) CipherSuite SSL_DH_RSA_WITH_3DES_EDE_CBC_SHA
            　　2.4.17) CipherSuite SSL_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA
            　　2.4.18) CipherSuite SSL_DHE_DSS_WITH_DES_CBC_SHA
            　　2.4.19) CipherSuite SSL_DHE_DSS_WITH_3DES_EDE_CBC_SHA
            　　2.4.20) CipherSuite SSL_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA
            　　2.4.21) CipherSuite SSL_DHE_RSA_WITH_DES_CBC_SHA
            　　2.4.22) CipherSuite SSL_DHE_RSA_WITH_3DES_EDE_CBC_SHA
            　　2.4.23) CipherSuite SSL_DH_anon_EXPORT_WITH_RC4_40_MD5
            　　2.4.24) CipherSuite SSL_DH_anon_WITH_RC4_128_MD5
            　　2.4.25) CipherSuite SSL_DH_anon_EXPORT_WITH_DES40_CBC_SHA
            　　2.4.26) CipherSuite SSL_DH_anon_WITH_DES_CBC_SHA
            　　2.4.27) CipherSuite SSL_DH_anon_WITH_3DES_EDE_CBC_SHA
            　　2.4.28) CipherSuite SSL_FORTEZZA_KEA_WITH_NULL_SHA
            　　2.4.29) CipherSuite SSL_FORTEZZA_KEA_WITH_FORTEZZA_CBC_SHA
            　　2.4.30) CipherSuite SSL_FORTEZZA_KEA_WITH_RC4_128_SHA
        　　2.5) 压缩方法
        3) server_hello:
        　　3.1) 版本(ProtocolVersion)
        　　代表服务端"采纳"的最高支持的SSL版本号
            　　3.1.1) 主版本: 3
            　　3.1.2) 次版本: 0
        　　3.2) 随机数(Random)
        　　服务端产生的一个用于生成主密钥(master key)的32字节的随机数(主密钥由客户端和服务端的随机数共同生成)
            　　3.2.1) uint32 gmt_unix_time;
            　　3.2.2) opaque random_bytes[28];
        　　4+28=32字节
        　　3.3) 会话ID: opaque SessionID<0..32>;
        　　3.4) 密文族(加密套件):
        　　代表服务端采纳的用于本次通讯的的加密套件
        　　3.5) 压缩方法:
        　　代表服务端采纳的用于本次通讯的的压缩方法
    　　总体来看，server_hello就是服务端对客户端的的回应，表示采纳某个方案
        4) certificate:
    　　SSL服务器将自己的"服务端公钥证书(注意，是公钥整数)"发送给SSL客户端
    　　ASN.1Cert certificate_list<1..2^24-1>;
        5) server_key_exchange:
        　　1) RSA
        　　执行RSA密钥协商过程
            　　1.1) RSA参数(ServerRSAParams)
                　　1.1.1) opaque RSA_modulus<1..2^16-1>;
                　　1.1.2) opaque RSA_exponent<1..2^16-1>;
           　　 1.2) 签名参数(Signature)
                　　1.2.1) anonymous: null
                　　1.2.2) rsa
                    　　1.2.2.1) opaque md5_hash[16];
                    　　1.2.2.2) opaque sha_hash[20];
                　　1.2.3) dsa
                    　　1.2.3.1) opaque sha_hash[20];
        　　2) diffie_hellman
        　　执行DH密钥协商过程
            　　2.1) DH参数(ServerDHParams)
                　　2.1.1) opaque DH_p<1..2^16-1>;
                　　2.1.2) opaque DH_g<1..2^16-1>;
                　　2.1.3) opaque DH_Ys<1..2^16-1>;
            　　2.2) 签名参数(Signature)
                　　2.2.1) anonymous: null
                　　2.2.2) rsa
                    　　2.2.2.1) opaque md5_hash[16];
                    　　2.2.2.2) opaque sha_hash[20];
                　　2.2.3) dsa
                    　　2.2.3.1) opaque sha_hash[20];
        　　3) fortezza_kea
        　　执行fortezza_kea密钥协商过程
            　　3.1) opaque r_s [128]
    6) certificate_request:
        6.1) 证书类型(CertificateType)
            6.1.1) RSA_sign
            6.1.2) DSS_sign
            6.1.3) RSA_fixed_DH
            6.1.4) DSS_fixed_DH
            6.1.5) RSA_ephemeral_DH
            6.1.6) DSS_ephemeral_DH
            6.1.7) FORTEZZA_MISSI
        6.2) 唯一名称(DistinguishedName)
        certificate_authorities<3..2^16-1>;
    7) server_done:
    服务器总是发送server_hello_done报文，指示服务器的hello阶段结束
    struct { } ServerHelloDone;
    8) certificate_verify:
    签名参数(Signature)
        8.1) anonymous: null
        8.2) rsa
            8.2.1) opaque md5_hash[16];
            8.2.2) opaque sha_hash[20];
        8.3) dsa
            8.3.1) opaque sha_hash[20];
    9) client_key_exchange:
        9.1) RSA
            9.1.1) PreMasterSecret
                9.1.1.1) ProtocolVersion
                9.1.1.2) opaque random[46];
        9.2) diffie_hellman: opaque DH_Yc<1..2^16-1>;
        9.3) fortezza_kea
            9.3.1) opaque y_c<0..128>;
            9.3.2) opaque r_c[128];
            9.3.3) opaque y_signature[40];
            9.3.4) opaque wrapped_client_write_key[12];
            9.3.5) opaque wrapped_server_write_key[12];
            9.3.6) opaque client_write_iv[24];
            9.3.7) opaque server_write_iv[24];
            9.3.8) opaque master_secret_iv[24];
            9.3.9) opaque encrypted_preMasterSecret[48];
    10) finished:
        10.1) opaque md5_hash[16];
        10.2) opaque sha_hash[20];

*/
package main

import (
	"fmt"
	"net"
)

func main() {
	Start()
}

func Start() {
	listener, err := net.Listen("tcp4", ":81")
	if err != nil {
		fmt.Println("监听网络错误:", err.Error())
	}
	for {
		buf := make([]byte, 1024)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go handler(buf[:n])
	}

}

func handler(buf []byte) {
	fmt.Println("客户端第一次请求内容：\n", buf, "\n")
	fmt.Println("记录号(22为“握手”记录)：", buf[0])
	fmt.Println("协议主版本号：", buf[1], "协议次版本号：", buf[2], "完整版本号：", buf[1], ".", buf[2])

}
