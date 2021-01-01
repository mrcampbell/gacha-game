extern crate fighter_service;
extern crate diesel;

use self::fighter_service::*;
use self::models::*;
use self::diesel::prelude::*;
use self::schema::units;


fn main() {
    use fighter_service::schema::units::dsl::*;

    let connection = establish_connection();

    diesel::insert_into(units)
    .values((id.eq("falcano"), name.eq("Falcano"), element.eq(1), type_.eq(4)))
    .execute(&connection);

    // let results = units
    //     .load::<Unit>(&connection)
    //     .expect("Error loading units");

    // println!("Displaying {} units", results.len());
    // for post in results {
    //     println!("{}", post.id);
    //     println!("----------\n");
    //     println!("{}", post.name);
    // }
}