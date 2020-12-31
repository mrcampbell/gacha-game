#[derive(Queryable)]
pub struct Unit {
    pub id: String,
    pub name: String,
    pub element: i32,
    pub type_: i32,
}
