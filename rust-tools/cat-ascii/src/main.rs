pub mod ascii;
pub mod cli;

use std::io;

use clap::{ArgMatches, Command};
use clap_complete::{generate, Generator, Shell};
use cli::build_cli;
use rand::{seq::SliceRandom, thread_rng};

use crate::ascii::CAT_ASCII;

fn main() {
    let match_result: ArgMatches = build_cli().get_matches();

    match match_result.subcommand() {
        Some(("list", _)) => list_names(),
        Some(("show", sub_matches)) => show_name(
            sub_matches
                .get_one::<String>("ASCII_NAME")
                .expect("required")
                .to_owned(),
        ),
        Some(("completion", sub_matches)) => {
            let shell = sub_matches
                .get_one::<Shell>("SHELL_NAME")
                .expect("required")
                .to_owned();
            let mut cmd = build_cli();
            eprintln!("Generating completion file for {shell}...");
            print_completions(shell, &mut cmd);
        }
        _ => pick_random_ascii(),
    }
}

fn list_names() {
    for ele in CAT_ASCII {
        println!("{}", ele.name)
    }
}

fn show_name(name: String) {
    println!("name {}:", name);
    for ele in CAT_ASCII {
        if ele.name == name {
            println!("{}", ele.ascii);
            println!("from: {}", ele.from);

            return;
        }
    }

    println!("Not found. Please use one of the following names:");
    list_names();
}

fn pick_random_ascii() {
    let mut rng = thread_rng();
    match CAT_ASCII.choose(&mut rng) {
        Some(cat_ascii) => println!("{}", cat_ascii.ascii),
        None => println!("No cat ascii is available"),
    }
}

fn print_completions<G: Generator>(gen: G, cmd: &mut Command) {
    generate(gen, cmd, cmd.get_name().to_string(), &mut io::stdout());
}
