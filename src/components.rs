use bevy::prelude::*;

#[derive(Component)]
pub struct Velocity {
    pub x: f32,
    pub y: f32,
}

#[derive(Component)]
pub struct BoxCollider {
    pub width: f32,
    pub height: f32,
}

#[derive(Component)]
pub struct Player;

#[derive(Component)]
pub struct Border;

#[derive(Component)]
pub struct Ground {
    pub last: bool,
}

#[derive(Component)]
pub enum AnimationState {
    Animating,
    UpdateTo(Animation),
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
pub enum Animation {
    PlayerFall,
    PlayerFly,
    PlayerDie,
}

#[derive(Component)]
pub struct Animator {
    pub timer: Timer,
    pub num_frames: usize,
}

impl Default for Animator {
    fn default() -> Self {
        Animator {
            timer: Timer::from_seconds(0.0, TimerMode::Once),
            num_frames: 0,
        }
    }
}
