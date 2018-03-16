/*
go语言中椭圆曲线加密算法的使用

椭圆曲线加密算法,使用golang的实现!
最近在看一些关于比特币的东西,里边有个椭圆曲线加密算法,查了下,感觉很不错!

与经典的RSA，DSA等公钥密码体制相比，椭圆密码体制有以下优点：
1.安全性高:有研究表示160位的椭圆密钥与1024位的RSA密钥安全性相同。
2.处理速度快:在私钥的加密解密速度上，ecc算法比RSA、DSA速度更快。
3.存储空间占用小。
4.带宽要求低.

椭圆曲线密码体制是目前已知的公钥体制中，对每比特所提供加密强度最高的一种体制。解椭圆曲线上的离散对数问题的最好算法是Pollard rho方法，
其时间复杂度为，是完全指数阶的。其中n为等式(2)中m的二进制表示的位数。当n=234, 约为2117，需要1.6x1023 MIPS 年的时间。而我们熟知的
RSA所利用的是大整数分解的困难问题，目前对于一般情况下的因数分解的最好算法的时间复杂度是子指数阶的，当n=2048时，需要2x1020MIPS年的时间。
也就是说当RSA的密钥使用2048位时，ECC的密钥使用234位所获得的安全强度还高出许多。它们之间的密钥长度却相差达9倍，当ECC的密钥更大时它们之间
差距将更大。更ECC密钥短的优点是非常明显的，随加密强度的提高，密钥长度变化不大。

第六届国际密码学会议对应用于公钥密码系统的加密算法推荐了两种：基于大整数因子分解问题（IFP）的RSA算法和基于椭圆曲线上离散对数计算问题
（ECDLP）的ECC算法。RSA算法的特点之一是数学原理简单、在工程应用中比较易于实现，但它的单位安全强度相对较低。目前用国际上公认的对于
RSA算法最有效的攻击方法--一般数域筛(NFS)方法去破译和攻击RSA算法，它的破译或求解难度是亚指数级的。ECC算法的数学理论非常深奥和复杂，
在工程应用中比较难于实现，但它的单位安全强度相对较高。用国际上公认的对于ECC算法最有效的攻击方法--Pollard rho方法去破译和攻击ECC算法，
它的破译或求解难度基本上是指数级的。正是由于RSA算法和ECC算法这一明显不同，使得ECC算法的单位安全强度高于RSA算法，也就是说，要达到同样的
安全强度，ECC算法所需的密钥长度远比RSA算法低（见表1和图1）。这就有效地解决了为了提高安全强度必须增加密钥长度所带来的工程实现难度的问题.

下面附带的代码,原本是用于网站用户密码的加密和校验的,当然也可以用于其他地方!
*/
package main

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var (
	runMode  string
	cfg      config.Configer //全局配置文件
	randKey  string
	randSign string
	prk      *ecdsa.PrivateKey
	puk      ecdsa.PublicKey
	curve    elliptic.Curve
)

func main() {
	//	example()
	//	reader := rand.Reader
	//	p, x, y, err := elliptic.GenerateKey(elliptic.P256(), reader)
	//	fmt.Println(p, x, y, err)

	//	fmt.Println(prk, puk)

	err := BuildKey()
	fmt.Println(err)

}

func BuildKey() error {

	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}
	ecder, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return err
	}
	keypem, err := os.OpenFile("ec-prk.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	err = pem.Encode(keypem, &pem.Block{Type: "EC PRIVATE KEY", Bytes: ecder})
	if err != nil {
		return err
	}

	//
	publicKey := &priv.PublicKey
	ecder, err = x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("111", err)
		return err
	}
	keypem, err = os.OpenFile("ec-puk.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	err = pem.Encode(keypem, &pem.Block{Type: "EC PUBLIC KEY", Bytes: ecder})
	if err != nil {
		return err
	}

	return nil

	//	prk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//	fmt.Println(prk, err)

	//	derStream := x509.MarshalECPrivateKey(prk)
	//	block := &pem.Block{
	//		Type:  "RSA PRIVATE KEY",
	//		Bytes: derStream,
	//	}
	//	file, err := os.Create("private.pem")
	//	if err != nil {
	//		return
	//	}
	//	err = pem.Encode(file, block)
	//	if err != nil {
	//		return
	//	}
}

func init() {
	var err error
	cfg, err = config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		return
	}
	randSign = cfg.String("RandSign")
	if len(randSign) == 0 {
		return
	}
	randKey = cfg.String("randKey")
	if len(randKey) == 0 {
		return
	}
	beego.Trace("Rand Key =", randKey)
	beego.Trace("Rand Sign =", randSign)
	//根据rand长度，使用相应的加密椭圆参数
	length := len([]byte(randKey))
	if length < 224/8 {
		beego.Error("The length of Rand Key is too small, Crypt init failed, Please reset it again !")
		return
	}
	if length >= 521/8+8 {
		beego.Notice("Rand length =", length, "Using 521 level !")
		curve = elliptic.P521()
	} else if length >= 384/8+8 {
		beego.Notice("Rand length =", length, "Using 384 level !")
		curve = elliptic.P384()
	} else if length >= 256/8+8 {
		beego.Notice("Rand length =", length, "Using 256 level !")
		curve = elliptic.P256()
	} else if length >= 224/8+8 {
		beego.Notice("Rand length =", length, "Using 244 level !")
		curve = elliptic.P224()
	}
	//创建密匙对
	prk, err = ecdsa.GenerateKey(curve, strings.NewReader(randKey))
	if err != nil {
		beego.Error("Crypt init fail,", err, " need = ", curve.Params().BitSize)
		return
	}
	puk = prk.PublicKey
}

//Encrypt 对Text进行加密，返回加密后的字节流
func Sign(text string) (string, error) {
	r, s, err := ecdsa.Sign(strings.NewReader(randSign), prk, []byte(text))
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(b.Bytes()), nil
}

//解密
func getSign(text, byterun []byte) (rint, sint big.Int, err error) {
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err != nil {
		err = errors.New("decode error," + err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	count, err := r.Read(buf)
	if err != nil {
		fmt.Println("decode =", err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]), "+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, " + err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, " + err.Error())
		return
	}
	return
}

//Verify 对密文和明文进行匹配校验
func Verify(text, passwd string) (bool, error) {
	byterun, err := hex.DecodeString(passwd)
	if err != nil {
		return false, err
	}
	rint, sint, err := getSign([]byte(text), byterun)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(&puk, []byte(text), &rint, &sint)
	return result, nil
}
