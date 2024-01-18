use std::{io, path::Path, fs};

use clap::{arg, command, value_parser, ArgAction, Command, ArgMatches, Arg};
use clap_complete::{Shell, Generator, generate};

const LINT_DIR_ARG_EMPTY: &str = "directory must have a non-whitespace value";
const LANG_ARG_EMPTY: &str    = "lang must have a non-whitespace value";

fn main() {
  let match_result: ArgMatches = build_cli().get_matches();

    match match_result.subcommand() {
        Some(("epub", sub_matches)) => {
          match sub_matches.subcommand() {
            Some(("compress-and-lint", sub_matches)) => {
              let lang = sub_matches
                .get_one::<String>("lang")
                .expect("required")
                .to_owned();

              let directory = sub_matches
                .get_one::<String>("directory")
                .expect("required")
                .to_owned();

              let compress_images = sub_matches
                .get_one::<bool>("compress-images")
                .expect("required")
                .to_owned();

              validate_compress_and_lint_args(directory.as_str(), lang.as_str());
              println!("lang: {} directory: {} images: {}", lang, directory, compress_images);

              folder_must_exist(directory.as_str());

              let epubs = must_get_all_files_with_ext_in_folder(&directory, "epub");
              for epub in epubs {
                println!("{}", epub)
              }

              unimplemented!()
          },
            _ => unimplemented!(),
          }
        },
        Some(("cbr", _)) => {
            println!("Not implemented...");
            unimplemented!()
        },
        Some(("cbz", _)) => {
          println!("Not implemented...");
          unimplemented!()
      },
        Some(("completion", sub_matches)) => {
            let shell = sub_matches
                .get_one::<Shell>("SHELL_NAME")
                .expect("required")
                .to_owned();
            let mut cmd = build_cli();
            eprintln!("Generating completion file for {shell}...");
            print_completions(shell, &mut cmd);
        }
        _ => unimplemented!(),
    }
}

fn build_cli() -> Command {
    command!()
        .about("A set of functions that are helpful for linting ebooks")
        .subcommand(
    Command::new("epub").about("Handles operations on epub files")
            .subcommand(Command::new("compress-and-lint")
            .arg(
              Arg::new("directory")
                .short('d')
                .long("directory")
                .default_value(".")
                .help("the location to run the epub lint logic")
            )
            .arg(
              Arg::new("lang")
                .short('l')
                .long("lang")
                .default_value("en")
                .help("the language to add to the xhtml, htm, or html files if the lang is not already specified")
            )
            .arg(
              Arg::new("compress-images")
                .short('i')
                .long("compress-images")
                .num_args(0)
                .help("whether or not to also compress images which requires imgp to be installed")
            )
          )
        )
        .subcommand(Command::new("cbr"))
        .subcommand(
          Command::new("cbz")
              // .about("Shows the specified cat ascii whose name is provided")
              // .arg(arg!(<ASCII_NAME> "The name of the cat ascii to display."))
              // .arg_required_else_help(true),
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

fn print_completions<G: Generator>(gen: G, cmd: &mut Command) {
  generate(gen, cmd, cmd.get_name().to_string(), &mut io::stdout());
}

fn validate_compress_and_lint_args(directory: &str, lang: &str) {
  assert_eq!(false, directory.trim().is_empty(), "{}", LINT_DIR_ARG_EMPTY);
  assert_eq!(false, lang.trim().is_empty(), "{}", LANG_ARG_EMPTY);
}

fn folder_must_exist(directory: &str) {
  let path = Path::new(directory);
  assert_eq!(true, path.is_dir(), "\"{}\" must be a folder", directory);
}

fn must_get_all_files_with_ext_in_folder(directory: &str, ext: &str) -> Vec<String> {
    let entries = fs::read_dir(directory).unwrap();
    let mut files_with_ext: Vec<String> = Vec::new();

    for entry in entries {
        let entry = entry.unwrap();
        let path = entry.path();

        if path.is_file() {
            let extension = path.extension().unwrap_or_default().to_str().unwrap();
            if extension == ext {
              files_with_ext.push(path.display().to_string())
            }
        }
    }

    return files_with_ext;
}