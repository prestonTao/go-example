c := http.Client{
        Transport: &http.Transport{
            Dial: func(netw, addr string) (net.Conn, error) {
                deadline := time.Now().Add(25 * time.Second)
                c, err := net.DialTimeout(netw, addr, time.Second*20)
                if err != nil {
                    return nil, err
                }
                c.SetDeadline(deadline)
                return c, nil
            },
        },
    }
 
//参考资料：
//http://www.reddit.com/r/golang/comments/121lc9/changing_httpget_timeout_and_not_having_too_many

//http://www.reddit.com/r/golang/comments/10awvj/timeout_on_httpget/