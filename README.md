# CLIAS
A simple CLI tool to add a new alias quickly to your zshrc file



### Build the Executable

In the terminal, navigate to the project directory (`updateZshrc`) and run the following command to build the executable:

```bash
go build
```

This command will create an executable file with the name of your project (`updateZshrc`).

###  Run the Executable

After the build is successful, you can run the executable:

```bash
./clias
```

This will prompt you to enter the alias name, alias value, and optionally the `zshrc` file path.

### Distribute or Install the Executable (Optional)

If you want to distribute or install the executable, you can move it to a directory in your `PATH` or distribute it along with the necessary instructions.

```bash
sudo mv clias /usr/local/bin
```

Now, you can run the `clias` command from anywhere in your terminal.

Remember to review the permissions and ensure that the directory containing your executable is in your `PATH` for convenient usage.
