# Nip17 Pipe

This command line tool can be used with NAK to test sending and receiving of NIP17 DMs over Nostr.

# Example

```

# sending
export NOSTR_PUBLIC_KEY=receiverpubkey
export NOSTR_SECRET_KEY=senderprivatekey

nip17pipe send -m "Hello Nip17" -r $NOSTR_PUBLIC_KEY |nak event --auth wss://example.relay.com

# receiving

export NOSTR_SECRET_KEY=receiverprivatekey
export NOSTR_PUBLIC_KEY=receiverpublickey

nak req --stream --auth -k 1059 -p $NOSTR_PUBLIC_KEY  --stream wss://example.relay.com | nip17pipe receive

```
