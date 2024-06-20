use crate::components::*;
use bevy::prelude::*;

const TILE_SIZE: f32 = 15.0;
const GAME_WIDTH: f32 = 300.0;
const GAME_HEIGHT: f32 = 600.0;

pub fn create_ground(
    commands: &mut Commands,
    texture_atlas_layouts: &mut ResMut<Assets<TextureAtlasLayout>>,
    asset_server: &Res<AssetServer>,
) {
    let window_left = -GAME_WIDTH / 2.0;
    let window_bottom = -GAME_HEIGHT / 2.0;
    let ground_y = window_bottom + TILE_SIZE / 2.0 * (1.5);

    commands.spawn((
        Transform::from_xyz(0.0, ground_y, 0.0),
        BoxCollider {
            width: GAME_WIDTH,
            height: TILE_SIZE,
        },
    ));

    println!("{}", ground_y);
    let num_ground = (GAME_WIDTH / TILE_SIZE).ceil() as isize;
    for i in 0..num_ground {
        // println!("{}", window_left + (0.5 + i as f32) * TILE_SIZE);
        create_ground_tile(
            commands,
            texture_atlas_layouts,
            asset_server,
            window_left + (0.5 + i as f32) * TILE_SIZE,
            ground_y,
        );
    }
}

fn create_ground_tile(
    commands: &mut Commands,
    texture_atlas_layouts: &mut ResMut<Assets<TextureAtlasLayout>>,
    asset_server: &Res<AssetServer>,
    x: f32,
    y: f32,
) {
    let texture_handle: Handle<Image> = asset_server.load("images/terrain.png");
    let texture_atlas = TextureAtlasLayout::from_grid(
        Vec2::new(TILE_SIZE, TILE_SIZE),
        4,
        2,
        None,
        Some(Vec2::new(0.0, 50.0)),
    );
    let texture_atlas_handle = texture_atlas_layouts.add(texture_atlas);

    commands.spawn((
        SpriteSheetBundle {
            texture: texture_handle,
            atlas: TextureAtlas {
                layout: texture_atlas_handle,
                index: 0,
            },
            transform: Transform::from_xyz(x, y, 0.0),
            ..Default::default()
        },
        Velocity { x: -50.0, y: 0.0 },
    ));
}
