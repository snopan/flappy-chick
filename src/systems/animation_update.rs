use crate::{components::*, Animations};
use bevy::prelude::*;

pub fn animation_update(
    animations: Res<Animations>,
    mut query: Query<(&Animation, &mut UpdateAnimation, &mut Animator, &mut Handle<Image>, &mut TextureAtlas)>
) {
    for (animation, mut update_animation, mut animator, mut texture, mut atlas) in &mut query {
        if !update_animation.0 { continue }

        let animation_info = match animations.0.get(animation) {
            Some(i) => i,
            None => panic!("no animation info found ")
        };

        animator.timer = Timer::from_seconds(animation_info.frame_duration, TimerMode::Repeating);
        animator.num_frames = animation_info.num_frames;

        *texture = animation_info.texture.clone();
        *atlas = animation_info.atlas.clone();

        update_animation.0 = false;
    }
}