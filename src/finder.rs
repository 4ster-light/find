use colored::*;
use std::fs::{self, ReadDir};
use std::path::{Path, PathBuf};

pub fn search_files(dir: &Path, filename: &str, show_dirs: &bool) -> Vec<PathBuf> {
    let mut results = Vec::new();

    match fs::read_dir(dir) {
        Ok(entries) => search_entries(entries, &mut results, filename, show_dirs),
        Err(e) => eprintln!("{}", format!("Error at {}: {}", dir.display(), e).red()),
    }

    results
}

fn search_entries(entries: ReadDir, results: &mut Vec<PathBuf>, filename: &str, show_dirs: &bool) {
    for entry in entries.filter_map(Result::ok) {
        let path = entry.path();

        if path.is_dir() {
            if *show_dirs {
                println!("{}", format!("Searching at {}", path.display()).green());
            }
            search_entries(fs::read_dir(&path).unwrap(), results, filename, show_dirs);
        } else if path.is_file() && path.file_name().and_then(|n| n.to_str()) == Some(filename) {
            results.push(path);
        }
    }
}
