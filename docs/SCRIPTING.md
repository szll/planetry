# Lua scripting

Everyone should be able to check for certain events e.g. if Earth moved out of the habitable zone or if Mars reaches a velocity/acceleration above a certain point without changing the application code itself.

To be able to do this, a Lua scripting VM is available. There are functions of the Go simulation scope available inside this VM. The functions can be used in order to get the information you need to write useful scripts.

The following code snipped shows a script, that is will check each simulation week (every 7th step), if the earth moved outside the habitable zone (either to close to or to far away from the sun).

```Lua
habitableZoneCheckDone = false

-- Both functions "getSteps" and "setPaused" as well as "AU" are exposed from the go context
function habitableZone()
  local steps = getSteps()
  if steps > 0 and steps % 7 == 0 and not habitableZoneCheckDone then
    local earth = getBodyById("earth")
    local sun = getBodyById("sun")
    local d = distance(earth.Position.X, earth.Position.Y, earth.Position.Z, sun.Position.X, sun.Position.Y, sun.Position.Z)  
    
    local toClose = d < AU * 0.95
    local toFar = d > AU * 2.4

    if toClose or toFar then
      if toClose then print("too close to the sun", d / AU, steps / 365) end
      if toFar then print("too far to the sun", d / AU, steps / 365) end
      setPaused(true)
      habitableZoneCheckDone = true
    end
  end
end

-- This function is also available to other scripts since it's in the global scope
distance = function (x1, y1, z1, x2, y2, z2)
  local dx = x2-x1
  local dy = y2-y1
  local dz = z2-z1
  return math.abs(math.sqrt(dx*dx+dy*dy+dz*dz))
end
```

> **NOTE**: Each script can only have one "main" function. The "main" function must be the only named function in the script file. If you have to use other functions, please write them as unnamed functions stored in vars, like "distance" in the example above.
The reason is: on file load, planetry will parse the file for one function of the notation "function \<name\>()" and stores this \<name\> in a map, which will be used in the simulation to call this script function. 

> **ATTENTION**: Most Go -> Lua functions will return refrences, so be careful when you alter values of bodies!

#### Functions and constants available in LuaVM

Currently available Go functions in Lua scope are:
 - `getBodyById(id string) *Body`: returns a body of the current scene by its id; the name is specified in the scene file
 - `getBodyByName(name string) *Body`: returns a body of the current scene by its name; the name is specified in the scene file
 - `getSteps() int`: returns the simulation steps (currently days)
 - `setPaused(paused bool)`: pauses or unpauses the simulation, depending on the `paused` value
 - `createPoint3D(x, y, z float64) *Point3D`: creates a new Point3D
 - `createVector3D(x, y, z float64) *Vector3D`: creates a new Vector3D
 - `createBody(name string, mass, radius float64, position *Point3D, velocity *Vector3D) *Body`: creates a new body`
 - `addBodyToScene(body *Body, red, green, blue, alpha int) err`: adds a new body to the scene

Also the constant value of the astronomical unit `AU` is available in the Lua scope.