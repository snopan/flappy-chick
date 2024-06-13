use bevy::ecs::component::Component;

#[derivce(Component)]
pub struct Position {
    pub x: usize,
    pub y: usize,
}

#[derivce(Component)]
pub struct PipeSize {
    pub width: usize,
    pub height: usize,
}

#[derive(Component)]
pub struct Player;