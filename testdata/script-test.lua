habitableZoneCheckDone = false

-- Both functions "getSteps" and "setPaused" as well as "AU" are exposed from the go context
function habitableZone()
  local steps = getSteps()
  if steps > 0 and steps % 7 == 0 and not habitableZoneCheckDone then
    local earth = getBodyById("earth")
    local sun = getBodyById("sun")

    if earth == nil then
      print("earth does not exist")
      habitableZoneCheckDone = true
      return
    end
    
    if sun == nil then
      print("sun does not exist")
      habitableZoneCheckDone = true
      return 
    end

    local d = distance(earth, sun)
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
distance = function (body1, body2)
  local dx = body2.Position.X - body1.Position.X
  local dy = body2.Position.Y - body1.Position.Y
  local dz = body2.Position.Z - body1.Position.Z
  return math.abs(math.sqrt(dx*dx+dy*dy+dz*dz))
end