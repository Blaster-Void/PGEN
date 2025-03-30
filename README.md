# PGEN
Random Character Generator This Go package generates random characters and provides functionalities to save them in a JSON file located in the user's home directory. The main purpose of the package is to create and manage passwords in a simple yet efficient way.
Features:

    Random Character Generation:

        By default, generates a random 8-character string.

        Users can customize the length of the string using the -c command-line flag.

    Password Storage:

        Saves generated passwords as a JSON file (mypassword.json) in the Documents folder.

    Password Retrieval:

        Displays previously saved passwords with the -show command-line flag.

Usage:

    Generate a Password: Run the program without arguments to generate an 8-character password. Use the -c [number] flag to specify the length of the password. Example: ./program -c 12 generates a 12-character password.

    View Saved Passwords: Use the -show flag to display the saved passwords from the JSON file. Example: ./program -show.

Implementation Details:

    The saveToJson function stores the password in JSON format.

    The ShowPasswords function retrieves and displays saved passwords.

    Characters are generated using ASCII codes between 65 and 122 (A-Z and additional characters).

    The program introduces a slight delay for decorative purposes during password generation.
    
    
# installing
-this tool helping you in install
-Important point : (go most install in your linux Device)
-`https://github.com/Blaster-Void/golanginstaller-bash` type `./installer main.go pgen` and done 
# tags
- Go

- Password-Manager

- CLI

- Random-Character-Generator

- JSON

- ASCII

- File-Handling

- Automation

- Programming

- Tool
