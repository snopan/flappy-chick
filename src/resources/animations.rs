use crate::components::*;
use bevy::{prelude::*, utils::*};

struct SpriteSheet<T: AsRef<str>> {
    path: T,
    tile_size: f32,
    columns: usize,
    rows: usize,
    frame_duration: f32,
    num_frames: usize,
}

const SPRITE_SHEETS: [(Animation, SpriteSheet<&'static str>); 3] = [
    (
        Animation::PlayerDie,
        SpriteSheet {
            path: "images/chicken_die.png",
            tile_size: 40.0,
            columns: 10,
            rows: 1,
            frame_duration: 0.1,
            num_frames: 10,
        },
    ),
    (
        Animation::PlayerFall,
        SpriteSheet {
            path: "images/chicken_fall.png",
            tile_size: 40.0,
            columns: 2,
            rows: 1,
            frame_duration: 0.1,
            num_frames: 2,
        },
    ),
    (
        Animation::PlayerFly,
        SpriteSheet {
            path: "images/chicken_fly.png",
            tile_size: 40.0,
            columns: 4,
            rows: 1,
            frame_duration: 0.1,
            num_frames: 4,
        },
    ),
];

#[derive(Resource)]
pub struct Animations(pub HashMap<Animation, AnimationInfo>);

pub struct AnimationInfo {
    pub texture: Handle<Image>,
    pub atlas: TextureAtlas,
    pub frame_duration: f32,
    pub num_frames: usize,
}

impl Default for Animations {
    fn default() -> Self {
        Animations(HashMap::new())
    }
}

impl Animations {
    pub fn load(
        &mut self,
        asset_server: &Res<AssetServer>,
        texture_atlas_layouts: &mut ResMut<Assets<TextureAtlasLayout>>,
    ) {
        for (animation, sheet) in SPRITE_SHEETS {
            let texture: Handle<Image> = asset_server.load(sheet.path);
            let layout = TextureAtlasLayout::from_grid(
                Vec2::new(sheet.tile_size, sheet.tile_size),
                sheet.columns,
                sheet.rows,
                None,
                None,
            );
            let texture_atlas_layout = texture_atlas_layouts.add(layout);
            self.0.insert(
                animation,
                AnimationInfo {
                    texture,
                    atlas: TextureAtlas {
                        layout: texture_atlas_layout,
                        index: 0,
                    },
                    frame_duration: sheet.frame_duration,
                    num_frames: sheet.num_frames,
                },
            );
        }
    }
}
