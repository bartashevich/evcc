<template>
	<div
		class="container container--loadpoint px-0 mb-md-2 d-flex flex-column justify-content-center"
	>
		<div
			ref="carousel"
			class="carousel d-lg-flex flex-wrap"
			:class="`carousel--${loadpoints.length}`"
		>
			<div
				v-for="(loadpoint, index) in loadpoints"
				:key="index"
				class="flex-grow-1 mb-3 m-lg-0 p-lg-0"
			>
				<Loadpoint
					v-bind="loadpoint"
					:id="index + 1"
					data-testid="loadpoint"
					:vehicles="vehicles"
					:smartCostType="smartCostType"
					:tariffGrid="tariffGrid"
					:tariffCo2="tariffCo2"
					:currency="currency"
					:multipleLoadpoints="loadpoints.length > 1"
					:gridConfigured="gridConfigured"
					:pvConfigured="pvConfigured"
					:batteryConfigured="batteryConfigured"
					class="h-100"
					:class="{ 'loadpoint-unselected': !selected(index) }"
					@click="scrollTo(index)"
				/>
			</div>
		</div>
		<div v-if="loadpoints.length > 1" class="d-flex d-lg-none justify-content-center">
			<button
				v-for="(loadpoint, index) in loadpoints"
				:key="index"
				class="btn btn-sm btn-link p-0 mx-1 indicator d-flex justify-content-center align-items-center evcc-default-text"
				:class="{ 'indicator--selected': selected(index) }"
				@click="scrollTo(index)"
			>
				<shopicon-filled-lightning
					v-if="isCharging(loadpoint)"
					class="indicator-icon"
				></shopicon-filled-lightning>
				<shopicon-filled-circle
					v-else-if="loadpoint.connected"
					class="indicator-icon"
				></shopicon-filled-circle>
				<shopicon-bold-circle v-else class="indicator-icon"></shopicon-bold-circle>
			</button>
		</div>
	</div>
</template>

<script>
import "@h2d2/shopicons/es/filled/circle";
import "@h2d2/shopicons/es/bold/circle";
import "@h2d2/shopicons/es/filled/lightning";

import Loadpoint from "./Loadpoint.vue";

export default {
	name: "Loadpoints",
	components: { Loadpoint },
	props: {
		loadpoints: Array,
		vehicles: Array,
		smartCostType: String,
		tariffGrid: Number,
		tariffCo2: Number,
		currency: String,
		gridConfigured: Boolean,
		pvConfigured: Boolean,
		batteryConfigured: Boolean,
	},
	data() {
		return { selectedIndex: 0, snapTimeout: null };
	},
	mounted() {
		this.$refs.carousel.addEventListener("scroll", this.handleCarouselScroll, false);
	},
	unmounted() {
		if (this.$refs.carousel) {
			this.$refs.carousel.removeEventListener("scroll", this.handleCarouselScroll);
		}
	},
	methods: {
		handleCarouselScroll() {
			const { scrollLeft } = this.$refs.carousel;
			const { offsetWidth } = this.$refs.carousel.children[0];
			this.selectedIndex = Math.round((scrollLeft - 7.5) / offsetWidth);
		},
		isCharging(lp) {
			return lp.charging && lp.chargePower > 0;
		},
		selected(index) {
			return this.selectedIndex === index;
		},
		scrollTo(index) {
			if (this.selectedIndex === index) {
				return;
			}
			this.selectedIndex = index;
			const $carousel = this.$refs.carousel;
			const width = $carousel.children[0].offsetWidth;
			$carousel.style.scrollSnapType = "none";
			$carousel.scrollTo({ top: 0, left: 7.5 + width * index, behavior: "smooth" });

			clearTimeout(this.snapTimeout);
			this.snapTimeout = setTimeout(() => {
				this.$refs.carousel.style.scrollSnapType = "x mandatory";
			}, 1000);
		},
	},
};
</script>
<style scoped>
.container--loadpoint {
	min-height: 300px;
}

@media (max-width: 991.98px) {
	.carousel {
		scroll-snap-type: x mandatory;
		overflow-x: scroll;
		display: flex;
		flex-wrap: nowrap !important;
		scrollbar-width: none; /* Firefox */
		-ms-overflow-style: none; /* IE 10+ */
	}
	.carousel::-webkit-scrollbar {
		display: none; /* Blink, Webkit */
	}
	.carousel > * {
		scroll-snap-align: center;
		min-width: 100%;
	}
	.indicator {
		width: 32px;
		height: 32px;
		opacity: 0.3;
		transition: opacity var(--evcc-transition-fast) ease-in;
	}
	.indicator--selected {
		opacity: 1;
	}
	.indicator-icon {
		width: 18px;
	}
	.loadpoint {
		opacity: 1;
		transform: scale(1);
		transition-property: opacity, transform;
		transition-duration: var(--evcc-transition-fast);
		transition-timing-function: ease-in;
	}
	.loadpoint-unselected {
		transform: scale(0.95);
		opacity: 0.5;
	}
}

/* show truncated tiles on breakpoint sm,md */
@media (min-width: 576px) and (max-width: 991.98px) {
	.container--loadpoint {
		max-width: none;
	}
	.carousel > *:first-child {
		margin-left: calc((100vw - var(--slide-width)) / 2);
	}
	.carousel > *:last-child {
		margin-right: calc((100vw - var(--slide-width)) / 2);
	}
	/* fixes safari issue with end-side padding https://webplatform.news/issues/2019-08-07 */
	.carousel::after {
		content: "";
		padding-right: 0.02px;
	}
	.carousel > * {
		min-width: var(--slide-width);
	}
}

/* breakpoint sm */
@media (min-width: 576px) and (max-width: 767.98px) {
	.carousel {
		--slide-width: 540px;
	}
}

/* breakpoint md */
@media (min-width: 768px) and (max-width: 991.98px) {
	.carousel {
		--slide-width: 720px;
	}
}

/* breakpoint lg, 2-col grid */
@media (min-width: 992px) {
	.carousel {
		display: grid !important;
		grid-gap: 2rem;
		grid-template-columns: repeat(auto-fit, minmax(450px, 1fr));
	}
}

/* breakpoint lg, tall screen, 2 loadpoints rows */
@media (min-width: 992px) and (min-height: 1450px) {
	.carousel--2 {
		grid-gap: 4rem;
		grid-template-columns: 1fr;
	}
}

/* breakpoint lg, taller screen, 3 loadpoints rows */
@media (min-width: 992px) and (min-height: 1900px) {
	.carousel--3 {
		grid-gap: 4rem;
		grid-template-columns: 1fr;
	}
}
</style>
