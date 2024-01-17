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

    if let Some(shell) = match_result.get_one::<Shell>("completion").copied() {
        let mut cmd = build_cli();
        eprintln!("Generating completion file for {shell}...");
        print_completions(shell, &mut cmd);

        return;
    }

    let is_list_action: bool = *match_result.get_one("list").unwrap_or(&false);
    if is_list_action {
        list_names();

        return;
    }

    let name = match_result.get_one::<String>("show");
    match name {
        Some(name) => show_name(name.to_owned()),
        None => pick_random_ascii(),
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
