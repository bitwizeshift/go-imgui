#include "imgui.h"

#include "shims.h"

extern "C" {

void shimTextColored(float r, float g, float b, float a, const char *s) {
	ImGui::TextColored(ImVec4(r, g, b, a), "%s", s);
}

void shimTextDisabled(const char *s) {
	ImGui::TextDisabled("%s", s);
}

void shimTextWrapped(const char *s) {
	ImGui::TextWrapped("%s", s);
}

void shimLabelText(const char *label, const char *s) {
	ImGui::LabelText(label, "%s", s);
}

void shimBulletText(const char *s) {
	ImGui::BulletText("%s", s);
}

void shimSetTooltip(const char *s) {
	ImGui::SetTooltip("%s", s);
}

} // extern "C"
