use clap::{arg, command, value_parser, ArgAction, Command};
use clap_complete::Shell;

pub fn build_cli() -> Command {
    command!()
        .about("Displays cat ascii when run.")
        .subcommand(Command::new("list").about("Lists names of all cat ascii options"))
        .subcommand(
            Command::new("show")
                .about("Shows the specified cat ascii whose name is provided")
                .arg(arg!(<ASCII_NAME> "The name of the cat ascii to display."))
                .arg_required_else_help(true),
        )
        .subcommand(
            Command::new("completion")
                .about("Generates code completion for the specified shell")
                .arg(
                    arg!(<SHELL_NAME> "The name of the shell to generate completion for.")
                        .action(ArgAction::Set)
                        .value_parser(value_parser!(Shell)),
                )
                .arg_required_else_help(true),
        )
}
