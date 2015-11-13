/******************************************************************************
 * Discordgo demo program.
 *
 * Please run this with command line arguments of email and password.
 */
package main

import (
	"fmt"
	"os"

	discord "github.com/bwmarrin/discordgo"
)

func main() {

	var err error
	var email string = os.Args[1]
	var password string = os.Args[2]

	// Create new session object and enable debugging.
	session := discord.Session{Debug: true}

	// Login to the Discord server with the provided  email and password
	// from the command line arguments
	session.Token, err = session.Login(email, password)
	if err != nil {
		fmt.Println("Unable to login to Discord.")
		fmt.Println(err)
		return
	}

	// Example using Request function to query a specific URL
	// This pulls authenticated user's information.
	// Request returns the actual request body not JSON
	body, err := discord.Request(&session, "http://discordapp.com/api/users/@me")
	fmt.Println(body)

	// Use the User function to do the same as above.  This function
	// returns a User structure
	user, err := discord.Users(&session, "@me")
	fmt.Println(user)

	// Use the PrivateChannels function to get a list of PrivateChannels
	// for a given user.
	private, err := discord.PrivateChannels(&session, "@me")
	fmt.Println(private)

	// Use the Servers function to pull all available servers for a given user
	// This returns a Server structure
	servers, err := discord.Servers(&session, "@me")
	fmt.Println(servers)

	// Use the Members function to pull all members of a given server.
	members, err := discord.Members(&session, servers[0].Id)
	fmt.Println(members)

	// Use the Channels function to pull all available channels for a given
	// server.  This returns a Channel structure.
	channels, err := discord.Channels(&session, servers[0].Id)
	fmt.Println(channels)

	// Use the Messages function to pull messages from the given channel
	// limit the result to 2 messages and don't provide beforeId or afterId
	messages, err := discord.Messages(&session, channels[0].Id, 2, 0, 0)
	fmt.Println(messages)

	// Use SendMessage to send a message to the given channel.
	responce, err := discord.SendMessage(&session, channels[0].Id, "Testing Discordgo")
	fmt.Println(responce)

	// Use the Logout function to Logout from the Discord server.
	discord.Logout(&session)
	return
}