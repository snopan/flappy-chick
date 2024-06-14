use bevy::prelude::*;

#[derive(Component)]
pub struct Position {
    pub x: usize,
    pub y: usize,
}

#[derive(Component)]
pub struct PipeSize {
    pub width: usize,
    pub height: usize,
}

#[derive(Component)]
pub struct Player;

#[derive(Component)]
pub enum AnimationState {
    Animating(Animation),
    ChangeTo(Animation),
}

pub enum Animation {
    PlayerFall,
    PlayerFly,
    PlayerDie,
}

#[derive(Component)]
pub struct Animator {
    pub timer: Timer,
    pub num_frames: usize
}
