# Demo puller:  pipeline Pull tcp://localhost:12345
# Demo pusher:  pipeline Push tcp://localhost:12345 "my message"
from go/github.com/gdamore/mangos/protocol import pull, push
from go/github.com/gdamore/mangos/transport import ipc, tcp

def Pull(url):
  sock = pull.NewSocket()
  sock.AddTransport(ipc.NewTransport())
  sock.AddTransport(tcp.NewTransport())
  sock.Listen(url)
  while True:
    msg = sock.Recv()
    print "Pull: RECEIVED %q" % msg 

def Push(url, msg):
  sock = push.NewSocket()
  sock.AddTransport(ipc.NewTransport())
  sock.AddTransport(tcp.NewTransport())
  sock.Dial(url)
  print "Push: SENDING %q" % msg
  sock.Send(msg)
  sock.Close()

def main(args):
  if len(args) == 2 and args[0] == "Pull":
    Pull(args[1])
  elif len(args) == 3 and args[0] == "Push":
    Push(args[1], args[2])
  else:
    raise 'Usage: pipeline Pull|Push URL [msg]'
