use clap::{command, value_parser, Arg, ArgAction, Command};
use clap_complete::Shell;

pub fn build_cli() -> Command {
    command!().about(
    "Displays cat ascii when run."
  ).arg(
    Arg::new("list")
      .short('l')
      .long("list")
      .help("Lists names of all cat ascii options")
      .num_args(0)
      .exclusive(true)
  ).arg(
    Arg::new("show")
      .short('s')
      .long("show")
      .help("Shows the specified cat ascii whose name is provided. Use the list command to see the list of available names.")
      .num_args(1)
      .exclusive(true)
  ).arg(
    Arg::new("completion")
        .long("completion")
        .help("Generates code completion for the specified shell")
        .action(ArgAction::Set)
        .exclusive(true)
        .value_parser(value_parser!(Shell)),
  )
}
