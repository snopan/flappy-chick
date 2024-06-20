use crate::components::*;
use bevy::prelude::*;
use bevy::window::*;

pub fn border(
    mut resize_reader: EventReader<WindowResized>,
    mut query: Query<&mut Style, With<Border>>,
) {
    for e in resize_reader.read() {
        for mut style in &mut query {
            let border_width = (e.width - 300.0) / 2.0;
            let border_height = (e.height - 600.0) / 2.0;

            style.border = UiRect {
                left: Val::Px(border_width),
                right: Val::Px(border_width),
                top: Val::Px(border_height),
                bottom: Val::Px(border_height),
            }
        }
    }
}
