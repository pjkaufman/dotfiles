use clap::{command, Arg, ArgMatches};
use rand::{seq::SliceRandom, thread_rng};

struct CatAscii {
    ascii: &'static str,
    name: &'static str,
    from: &'static str,
}

const CAT_ASCII: [CatAscii; 8] = [
    CatAscii {
        ascii: "
  ───────────────────────────────────────
  ───▐▀▄───────▄▀▌───▄▄▄▄▄▄▄─────────────
  ───▌▒▒▀▄▄▄▄▄▀▒▒▐▄▀▀▒██▒██▒▀▀▄──────────
  ──▐▒▒▒▒▀▒▀▒▀▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▀▄────────
  ──▌▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▄▒▒▒▒▒▒▒▒▒▒▒▒▀▄──────
  ▀█▒▒▒█▌▒▒█▒▒▐█▒▒▒▀▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▌─────
  ▀▌▒▒▒▒▒▒▀▒▀▒▒▒▒▒▒▀▀▒▒▒▒▒▒▒▒▒▒▒▒▒▒▐───▄▄
  ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▌▄█▒█
  ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒█▒█▀─
  ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒█▀───
  ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▌────
  ─▌▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▐─────
  ─▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▌─────
  ──▌▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▐──────
  ──▐▄▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▄▌──────
  ────▀▄▄▀▀▀▀▀▄▄▀▀▀▀▀▀▀▄▄▀▀▀▀▀▄▄▀────────
    
  ",
        name: "big-cat",
        from: "https://www.asciiartcopy.com/ascii-cat.html",
    },
    CatAscii {
        ascii: r#"
                    .............                .""".             .""".
            ..."""""             """""...       $   . ".         ." .   $
        ..""        .   .   .   .   .    ..    $   $$$. ". ... ." .$$$   $
      ."    . " . " . " . " . " . " . " .  "" ."  $$$"""  "   "  """$$$  ".
    ."      . " . " . " . " . " . " . " .     $  "                    "   $
   ."   . " . " . "           "   " . " . "  ."      ...          ...     ".
  ."    . " . "    .."""""""""...     " . "  $     .$"              "$.    $
 ."     . " . " .""     .   .    ""..   . " $ ".      .""$     .""$      ." $
."    " . " .       . " . " . " .    $    " $ "      "  $$    "  $$       " $
$     " . " . " . " . " . " . " . "   $     $             $$.$$             $
$     " . " . " . " . " . " . " . " .  $  " $  " .        $$$$$        . "  $
$     " . " . " . " . " . " . " . " .  $    $      "  ..   "$"   ..  "      $
".    " . " . " . " . " . " . " . "   ."  "  $  . . . $  . .". .  $ . . .  $
 $    " . " . " . " . " . " . " . "  ."   "            ".."   ".."
  $     . " . " . " . " . " . "   .."   . " . "..    "             "    .."
  ".      " . " . " . " . " .  .""    " . " .    """$...         ...$"""
   ". "..     " . " . " . " .  "........  "    .....  ."""....."""
     ". ."$".....                       $..."$"$"."   $".$"... `":....
       "".."    $"$"$"$"""$........$"$"$"  ."."."  ...""      ."".    `"".
           """.$.$." ."  ."."."    ."."." $.$.$"""".......  ". ". $ ". ". $
                  """.$.$.$.$.....$.$.""""               ""..$..$."..$..$."

"#,
        name: "big-happy-cat",
        from: "https://ascii.co.uk/art/cats",
    },
    CatAscii {
        ascii: r#"
                                     ,
              ,-.       _,---._ __  / \
             /  )    .-'       `./ /   \
            (  (   ,'            `/    /|
             \  `-"             \'\   / |
              `.              ,  \ \ /  |
               /`.          ,'-`----Y   |
              (            ;        |   '
              |  ,-.    ,-'         |  /
              |  | (   |        hjw | /
              )  |  \  `.___________|/
              `--'   `--'
"#,
        name: "box-search-cat",
        from: "https://ascii.co.uk/art/cats",
    },
    CatAscii {
        ascii: r#"
      Art by Blazej Kozlowski
             _                        
             \`*-.                    
              )  _`-.                 
             .  : `. .                
             : _   '  \               
             ; *` _.   `*-._          
             `-.-'          `-.       
               ;       `       `.     
               :.       .        \    
               . \  .   :   .-'   .   
               '  `+.;  ;  '      :   
               :  '  |    ;       ;-. 
               ; '   : :`-:     _.`* ;
      [bug] .*' /  .*' ; .*`- +'  `*' 
            `*-*   `*-*  `*-*'
      "#,
        name: "bug-cat",
        from: "https://www.asciiart.eu/animals/cats",
    },
    CatAscii {
        ascii: r#"
                                _.---.
                      |\---/|  / ) ca|
          ------------;     |-/ /|foo|---
                      )     (' / `---'
          ===========(       ,'==========
          ||   _     |      |
          || o/ )    |      | o
          || ( (    /       ;
          ||  \ `._/       /
          ||   `._        /|
          ||      |\    _/||
        __||_____.' )  |__||____________
         ________\  |  |_________________
                  \ \  `-.
                   `-`---'  hjw
"#,
        name: "cat-food-cat",
        from: "https://user.xmission.com/~emailbox/ascii_cats.htm",
    },
    CatAscii {
        ascii: r#"
_____$$_____$$
_____$$$___$$$
_____$$$$$$$$$_______$$
______$$$$$$$_______$$$
_______$$$$$________$$$$
________$$$__________$$$$
________$$$$__________$$$
________$$$$$$$______$$$
________$$$$$$$$____$$$
_________$$$$$$$____$$
__________$$$$$$___$$
__________$$$$$_$__$$
__________$$$$_$$$_$$
__________$$$_$$$$$_$$
_________$$$__$$$$_$$
"#,
        name: "dollar-cat",
        from: "https://www.asciiartcopy.com/ascii-cat.html",
    },
    CatAscii {
        ascii: r#"
_._     _,-'""`-._
(,-.`._,'(       |\`-/|
    `-.-' \ )-`( , o o)
          `-    \`_`"'-
"#,
        name: "scared-cat",
        from: "https://www.asciiartcopy.com/ascii-cat.html",
    },
    CatAscii {
        name: "yarn-and-cat",
        ascii: "
            .-o=o-.
        ,  /=o=o=o=\\ .--.
       _|\\|=o=O=o=O=|    \\
   __.'  a`\\=o=o=o=(`\\   /
   '.   a 4/`|.-\"\"'`\\ \\ ;'`)   .---.
     \\   .'  /   .--'  |_.'   / .-._)
      `)  _.'   /     /`-.__.' /
   jgs `'-.____;     /'-.___.-'
                `\"\"\"`
",
        from: " https://user.xmission.com/~emailbox/ascii_cats.htm",
    },
];

fn main() {
    let match_result: ArgMatches = command!().about(
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
  ).get_matches();

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
