# Nip17 Pipe

This command line tool can be used with NAK to test sending and receiving of NIP17 DMs over Nostr.

# Requirements for example

* [nak command line tool](https://github.com/nbd-wtf/nak)

# Example

Note that the `nak` command line tool uses the `NOSTR_SECRET_KEY` environment variable to authenticate with the relay.

Sending
```
export NOSTR_PUBLIC_KEY=pubkey of RECEIVER
export NOSTR_SECRET_KEY=SECRETkey of SENDER

nip17pipe send -m "Hello Nip17" -r $NOSTR_PUBLIC_KEY |nak event --auth wss://example.relay.com
```

Receiving
```

export NOSTR_SECRET_KEY=pubkey of RECEIVER
export NOSTR_PUBLIC_KEY=pubkey of RECEIVER

nak req --stream --auth -k 1059 -p $NOSTR_PUBLIC_KEY --stream wss://example.relay.com | nip17pipe receive

```