use colored::*;
use std::fs;
use std::path::{Path, PathBuf};

pub fn search_files(dir: &Path, filename: &str, show_dirs: &bool) -> Vec<PathBuf> {
    let mut results = Vec::new();

    match fs::read_dir(dir) {
        Ok(entries) => {
            for entry in entries.filter_map(Result::ok) {
                let path = entry.path();

                if path.is_dir() {
                    if *show_dirs {
                        println!("{}", format!("Searching at {}", path.display()).green());
                    }
                    results.extend(search_files(&path, filename, show_dirs));
                } else if path.is_file() {
                    if path.file_name().and_then(|n| n.to_str()) == Some(filename) {
                        results.push(path);
                    }
                }
            }
        }
        Err(e) => {
            eprintln!(
                "{}",
                format!("Error reading {}: {}", dir.display(), e).red()
            );
        }
    }

    results
}
