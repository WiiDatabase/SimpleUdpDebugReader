package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	address    = "0.0.0.0:4405"
	bufferSize = 4096
	logFile    = "GeckoLog.txt"
)

func main() {
	// Open log file
	logFile, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Setup UDP connection
	conn, err := net.ListenPacket("udp4", address)
	if err != nil {
		log.Fatalf("Error setting up UDP connection: %v", err)
	}
	defer conn.Close()

	// Create channel to control goroutines
	done := make(chan bool)

	fmt.Println("Welcome to Simple UDP Debug Reader!")
	fmt.Println("")
	fmt.Println("Log is automatically saved to \"GeckoLog.txt\"")
	fmt.Println("c - Clear console")
	fmt.Println("q - Exit")
	fmt.Println("")

	// Handle incoming packets
	go func() {
		buf := make([]byte, bufferSize)
		for {
			n, _, err := conn.ReadFrom(buf)
			if err != nil {
				log.Printf("Error reading from UDP: %v", err)
				continue
			}

			data := buf[:n]
			msg := fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339), string(data))
			fmt.Print(msg)

			if _, err := logFile.WriteString(msg); err != nil {
				log.Printf("Error writing to log file: %v", err)
			}
		}
	}()

	// Handle keyboard input
	go func() {
		if err := keyboard.Open(); err != nil {
			log.Fatalf("Error opening keyboard: %v", err)
		}
		defer keyboard.Close()

		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				log.Printf("Error reading keyboard input: %v", err)
				continue
			}

			if key == keyboard.KeyCtrlC || char == 'q' {
				done <- true
				break
			} else if char == 'c' {
				print("\033[H\033[2J")
			}
		}
	}()

	// Wait for exit signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sig:
	case <-done:
	}

	fmt.Println("Exiting...")
}
