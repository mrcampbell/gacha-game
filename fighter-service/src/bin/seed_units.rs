extern crate fighter_service;
extern crate diesel;

use self::fighter_service::*;
use self::diesel::prelude::*;


fn main() {
    use fighter_service::schema::units::dsl::*;

    let connection = establish_connection();

    /*
    enum UnitType {
        UNIT_TYPE_UNKNOWN = 0;
        UNIT_TYPE_TANK = 1;
        UNIT_TYPE_HEALER = 2;
        UNIT_TYPE_SUPPORT = 3;
        UNIT_TYPE_ATTACKER = 4;
        }

    enum Element {
        ELEMENT_UNKNOWN = 0;
        ELEMENT_FIRE = 1;
        ELEMENT_WATER = 2;
        ELEMENT_GRASS = 3;
        ELEMENT_NORMAL = 4;
        }
 */

    diesel::insert_into(units)
    .values(
        &vec![
            // (id.eq("falcano"), name.eq("Falcano"), element.eq(1), unit_type.eq(4)),
            // (id.eq("fridgety"), name.eq("Fridgety"), element.eq(2), unit_type.eq(4)),
            // (id.eq("fightcus"), name.eq("Fightcus"), element.eq(3), unit_type.eq(4)),
            // (id.eq("talon"), name.eq("Talon"), element.eq(4), unit_type.eq(4)),

            (id.eq("1"), name.eq("Fire Tank"), element.eq(1), unit_type.eq(1)),
            (id.eq("2"), name.eq("Water Tank"), element.eq(2), unit_type.eq(1)),
            (id.eq("3"), name.eq("Grass Tank"), element.eq(3), unit_type.eq(1)),
            (id.eq("4"), name.eq("Normal Tank"), element.eq(4), unit_type.eq(1)),

            (id.eq("5"), name.eq("Fire Healer"), element.eq(1), unit_type.eq(2)),
            (id.eq("6"), name.eq("Water Healer"), element.eq(2), unit_type.eq(2)),
            (id.eq("7"), name.eq("Grass Healer"), element.eq(3), unit_type.eq(2)),
            (id.eq("8"), name.eq("Normal Healer"), element.eq(4), unit_type.eq(2)),

            (id.eq("9"), name.eq("Fire Support"), element.eq(1), unit_type.eq(3)),
            (id.eq("10"), name.eq("Water Support"), element.eq(2), unit_type.eq(3)),
            (id.eq("11"), name.eq("Grass Support"), element.eq(3), unit_type.eq(3)),
            (id.eq("12"), name.eq("Normal Support"), element.eq(4), unit_type.eq(3)),

            (id.eq("13"), name.eq("Fire Attacker"), element.eq(1), unit_type.eq(4)),
            (id.eq("14"), name.eq("Water Attacker"), element.eq(2), unit_type.eq(4)),
            (id.eq("15"), name.eq("Grass Attacker"), element.eq(3), unit_type.eq(4)),
            (id.eq("16"), name.eq("Normal Attacker"), element.eq(4), unit_type.eq(4)),

            ]
    )
    .execute(&connection)
    .expect("failed to insert units");
}