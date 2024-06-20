use crate::components::*;
use crate::resources::*;
use bevy::prelude::*;

pub fn animation_update(
    animations: Res<Animations>,
    mut query: Query<(
        &mut AnimationState,
        &mut Animator,
        &mut Handle<Image>,
        &mut TextureAtlas,
    )>,
) {
    for (mut animation_state, mut animator, mut texture, mut atlas) in &mut query {
        match *animation_state {
            AnimationState::UpdateTo(animation) => {
                let animation_info = match animations.0.get(&animation) {
                    Some(i) => i,
                    None => panic!("no animation info found "),
                };

                animator.timer =
                    Timer::from_seconds(animation_info.frame_duration, TimerMode::Repeating);
                animator.num_frames = animation_info.num_frames;

                *texture = animation_info.texture.clone();
                *atlas = animation_info.atlas.clone();

                *animation_state = AnimationState::Animating
            }
            AnimationState::Animating => {}
        }
    }
}
