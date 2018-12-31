# Describing the simulated scene

Planetry needs data about what it's going to simulate, e.g. the solarsystem. A scene file is this kind of description. The following snipped shows the structure of a scene file:

```json
{
  "version": 1,
  "meta": {
    "units": {
      "position": "au",
      "velocity": "km/s",
      "radius": "km"
    }
  },
  "backgroundColor": {
    "red": 22,
    "green": 22,
    "blue": 22,
    "alpha": 250
  },
  "scripts": [
    "script-sun-disappears-after-365-steps.lua"
  ],
  "bodies": [
    {
      "id": "sun",
      "name": "Sun",
      "mass": 1.98892e30,
      "radius": 1.0,
      "position": {
        "x": 0.0,
        "y": 0.0,
        "z": 0.0
      },
      "velocity": {
        "x": 0.0,
        "y": 0.0,
        "z": 0.0
      },
      "color": {
        "red": 255,
        "green": 255,
        "blue": 60,
        "alpha": 250
      }
    },
    ...
  ]
}
```

It consists of the following properties:
 - **version** (int): is the version of the scene file structure.
 - **meta** (object): meta itself is an object which contains meta data about the scene.
   - **units** (object): contains information about the units used for the **bodies** in the current scene file.
     - **position** (string): unit used for position values.
     - **velocity** (string): unit used for velocity values.
     - **radius** (string): unit used for radius value.
 - **backgroundColor** (object): defines the scene background color and has the typical color properties: **red**, **green**, **blue**, **alpha**.
 - **targetId** (string, *optional*): is the id of a body; this body will be centered within the view.
 - **scripts** (array of strings, *optional*): list of [scripts](SCRIPTING.md) that should get executed during simulation.
 - **bodies** (array of objects): this is the representation of all the objects in the scene like stars or planets.

<!--
  TODO: describe bodies
-->

Example: see `testdata/system.json`