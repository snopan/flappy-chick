use bevy::prelude::*;

#[derive(Component)]
pub struct Velocity {
    pub x: f32,
    pub y: f32,
}

#[derive(Component)]
pub struct PipeSize {
    pub width: usize,
    pub height: usize,
}

#[derive(Component)]
pub struct Player;

#[derive(Component, Clone, Copy, PartialEq, Eq, Hash)]
pub enum Animation {
    PlayerFall,
    PlayerFly,
    PlayerDie,
}

#[derive(Component)]
pub struct UpdateAnimation(pub bool);

#[derive(Component)]
pub struct Animator {
    pub timer: Timer,
    pub num_frames: usize
}

impl Default for Animator {
    fn default() -> Self {
        Animator{
            timer: Timer::from_seconds(0.0, TimerMode::Once),
            num_frames: 0,
        }
    }
}
