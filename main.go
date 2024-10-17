package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jeremyd/crusher17"
	"github.com/nbd-wtf/go-nostr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message",
	Long:  `Send a message using the provided private key and the recipient's public key.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Load sender's private key from environment variable
		sk := viper.GetString("NOSTR_SECRET_KEY")
		if sk == "" {
			fmt.Println("NOSTR_SECRET_KEY not set in the environment")
			os.Exit(1)
		}

		message, _ := cmd.Flags().GetString("message")
		relay, _ := cmd.Flags().GetString("relay")
		receivers, _ := cmd.Flags().GetStringArray("receiver")
		fmt.Println("Sending message to", receivers)

		receiverMap := make(map[string]string)
		for _, pk := range receivers {
			receiverMap[pk] = relay
		}

		gwe := &crusher17.GiftWrapEvent{
			SenderSecretKey: sk,
			SenderRelay:     relay,
			ReceiverPubkeys: receiverMap,
			Content:         message,
			Subject:         "",
			GiftWraps:       make(map[string]string),
		}

		// Wrap the message
		err := gwe.Wrap()
		if err != nil {
			fmt.Printf("Error wrapping message: %v\n", err)
			return
		}

		// Print the gift wraps for each receiver
		for _, giftWrap := range gwe.GiftWraps {
			fmt.Printf("%s\n", giftWrap)
		}
	},
}

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive a message",
	Long:  `Receive a message using the provided private key and the gift.`,
	Run: func(cmd *cobra.Command, args []string) {

		sk := viper.GetString("NOSTR_SECRET_KEY")
		if sk == "" {
			fmt.Println("NOSTR_SECRET_KEY not set in the environment")
			os.Exit(1)
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			gift := strings.TrimSpace(scanner.Text())
			// Your logic here
			received, err := crusher17.ReceiveMessage(sk, gift)

			if err != nil {
				fmt.Println("error receiving message: ", err)
			}

			var ev nostr.Event

			jsonerr := json.Unmarshal([]byte(received), &ev)
			if jsonerr != nil {
				fmt.Println("error decoding message: ", jsonerr)
			}

			fmt.Printf("Message from %s: %s\n", ev.PubKey, ev.Content)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(receiveCmd)

	sendCmd.Flags().StringArrayP("receiver", "r", []string{}, "Receiver public key (can be used multiple times)")
	sendCmd.Flags().StringP("message", "m", "", "Message to send")
	sendCmd.Flags().StringP("relay", "R", "wss://auth.nostr1.com", "relay to use")

	sendCmd.MarkFlagRequired("message")
	sendCmd.MarkFlagRequired("receiver")

	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
