package scan_test

import (
	"net"
	"strconv"
	"testing"

	"github.com/gopheramit/pScan/scan"
)

func TestStateString(t *testing.T) {
	ps := scan.PortState{}
	if ps.Open.String() != "closed" {
		t.Errorf("Expected %q,got%qinstead\n", ps.Open.String())
	}
	ps.Open = true
	if ps.Open.String() != "open" {
		t.Errorf("Expected%q,got%q instead\n", ps.Open.String())
	}
}

func TestRunHostFound(t *testing.T){
	testCases:=[]struct{
		name string
		expectState string
	}
	{
		{"Opne port","open"},
		{
			"ClosedPort","closed"
		}
	}
	host:="localhost"
	hl:=&scan.HostList{}
	hl.Add(host)

	ports:=[]int{}
	for _,tc:=range testCases{
		ln,err:=net.Listen("tcp",net.JoinHostPort(host,"0"))
		if err!=nil{
			t.Fatal(err)
		}
		defer ln.Close()

		_,portStr,err:=net.SplitHostPort(ln.Addr().String())
		if err!=nil{
			t.Fatal(err)
		}
		port,err:=strconv.Atoi(portStr)
		if err!=nil{
			t.Fatal(err)
		}
		ports=append(ports, port)
		if tc.name=="ClosedPort"{
			ln.Close()
		}

		res:=scan.Run(hl,ports)
		if len(res)!=1{
			t.Fatalf("Expected 1 result,got%d instesad\n",len(res))
		}
		if res[0].Host!=host{
			t.Errorf("Expected host%q,got%qinstead\n",host,res[0].Host)
		}
		if res[0].NotFound{
			t.Errorf("Expected host%q,to be found\n",host)
		}

		if len(res[0].PortState!=2){
			t.Fatalf("Expected 2 port states,got %d instean",len(res[0].PortState))
		}
		for i,tc:=range testCases{
			if res[0].PortState[i].Port!=port[i]{
				t.Errorf("Expected port %d,got%dinstead\n",port[0],res[0].PortState[i].Port)
			}
			if res[0].PortState[i].Open.String()!=tx.ecpectState{
				t.Errorf("Expected port %d to be %s\n",ports[i],tc.expectState)
			}
		}
	}

}

func TestRunHostNotFound(t *testing.T){
	host:="389.389.389.389"
	hl:=&scan.HostList{}
	hl.Add(host)
	res:=scan.Run(hl,[]int{})
	if len(res)!=1{
		t.Fatalf("Expected 1 results,got,%dinstead\n",len(res))
	}
	if res[0].Host!=host{
		t.Errorf("Expected host%q,got %q instead\n",host,res[0].Host)
	}
	if !res[0].NotFound{
		t.Errorf("Expected host %q not be found in ",host)
	}
	if len(res[0].PortState)!=0{
		t.Fatalf("Expected 0 port states ,got %d instead\n",len(res[0].PortState))
	}
}