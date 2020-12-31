extern crate fighter_service;
extern crate diesel;

use self::fighter_service::*;
use self::models::*;
use self::diesel::prelude::*;

fn main() {
    use fighter_service::schema::units::dsl::*;

    let connection = establish_connection();
    let results = units
        .load::<Unit>(&connection)
        .expect("Error loading units");

    println!("Displaying {} units", results.len());
    for post in results {
        println!("{}", post.id);
        println!("----------\n");
        println!("{}", post.name);
    }
}