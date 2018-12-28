habitableZoneCheckDone = false

-- Both functions "getSteps" and "setPaused" as well as "AU" are exposed from the go context
function habitableZone()
  local steps = getSteps()
  if steps > 0 and steps % 7 == 0 and not habitableZoneCheckDone then
    local earth = getBodyByName("Earth")
    local sun = getBodyByName("Sun")
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