local letSunDisappearDone = false
function letSunDisappear()
  if getSteps() >= 365 and not letSunDisappearDone then
    if getBodyById("sun") ~= nil then
      removeBodyById("sun")
      letSunDisappearDone = true
    end
  end
end
