employee-records-go
-------

Allows user to list / show employees

#### 1. Usage

  - List Employees

    `employees list`

  - Show employee

    `employees show ID`

#### 2. Building
- ##### Prerequisites
  - Install go

    ```
      $ git clone https://github.com/syndbg/goenv.git ~/.goenv```
    ```
    $ echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.bash_profile
    $ echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.bash_profile
    ```
  - Setup your GOPATH
  - Clone repo into your gopath
- ##### Run or build

    ```go run cmd/employees <args>```

    ```go build cmd/employees```

#### 3. Testing
    make test
