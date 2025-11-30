# padel-schedule-manager
Tool to notify when there is padel field availability for particular a date time / to book / to keep a slot unavailable

The idea is to be able to define a desire date time (could be recurrent) and this tool should be able to do some actions regarding the scheduler:

- books as soon as it becomes available (you need to have credits in your account)
- notifies you when the desire schedule becomes available. No further actions
- starts the booking process keeping infinitely the time slot unavailable (in case you want to book for two hours, but pay for one.)

## How to use it

### Run with default settings (notify mode)
```bash
make run
```

### Run in interactive mode
```bash
make run-interactive
```

### Run with specific execution modes
```bash
make run-reserve   # Reserve mode
make run-notify    # Notify mode
make run-freeze    # Freeze slot mode
```

### Use custom config file
```bash
make run-custom CONFIG=.env.production
```

### Combine parameters
```bash
make run-full MODE=reserve CONFIG=.env.custom
```

### Build the application
```bash
make build
./padel-schedule-manager -mode=reserve -config=.env.local
```

### Show all available commands
```bash
make help
```
