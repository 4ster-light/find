use std::path::Path;
use std::time::Instant;

use clap::Parser;
use colored::Colorize;

mod finder;

/// Find files by name in a specified directory
#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {
    /// The name of the file to search for
    #[arg(short, long)]
    filename: String,

    /// The directory to search
    #[arg(short, long, default_value = ".")]
    dir: String,

    /// Show the time elapsed for the search
    #[arg(short, long)]
    time: bool,

    /// Show the directories being searched at the moment
    #[arg(short, long)]
    show_dirs: bool,
}

fn main() {
    let args = Args::parse();
    let (
        filename,
        dir,
        time,
        show_dirs
    ) = (&args.filename, &args.dir, &args.time, &args.show_dirs);

    let start = Instant::now();
    let results = finder::search_files(&Path::new(dir), filename, show_dirs);
    let elapsed = start.elapsed();

    println!("");

    if *time {
        println!("{}", format!("Time elapsed: {:?}", elapsed).yellow().bold());
    }

    for result in results {
        println!(
            "{}",
            format!(
                "Found an occurrence of {} at {}",
                filename,
                result.display()
            )
            .blue()
            .bold()
        );
    }
}
