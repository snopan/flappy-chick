use crate::components::*;
use bevy::prelude::*;

pub fn create_border(commands: &mut Commands, window: &Window) {
    let border_width = (window.width() - 300.0) / 2.0;
    let border_height = (window.height() - 600.0) / 2.0;

    commands.spawn((
        NodeBundle {
            style: Style {
                width: Val::Vw(100.0),
                height: Val::Vh(100.0),
                border: UiRect {
                    left: Val::Px(border_width),
                    right: Val::Px(border_width),
                    top: Val::Px(border_height),
                    bottom: Val::Px(border_height),
                },
                align_items: AlignItems::Center,
                justify_content: JustifyContent::Center,
                ..Default::default()
            },
            border_color: Color::BLACK.into(),
            ..Default::default()
        },
        Border,
    ));
}
