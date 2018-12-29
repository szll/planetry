local letSunDisappearDone = false
function letSunDisappear()
  if getSteps() >= 365 and not letSunDisappearDone then
    local sun = getBodyById("sun")
    if sun ~= nil then
      removeBodyById("sun")
      letSunDisappearDone = true
    end
  end
end
