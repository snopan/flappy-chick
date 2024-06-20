use crate::components::*;
use bevy::prelude::*;

pub fn animate_sprite(time: Res<Time>, mut query: Query<(&mut Animator, &mut TextureAtlas)>) {
    for (mut animator, mut atlas) in &mut query {
        if animator.num_frames == 0 {
            return;
        }

        animator.timer.tick(time.delta());
        if animator.timer.just_finished() {
            atlas.index = if atlas.index == animator.num_frames - 1 {
                0
            } else {
                atlas.index + 1
            };
        }
    }
}
